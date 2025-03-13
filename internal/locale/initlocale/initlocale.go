// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package initlocale

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/divVerent/aaaaxy/internal/exitstatus"
	"github.com/divVerent/aaaaxy/internal/flag"
	"github.com/divVerent/aaaaxy/internal/font"
	"github.com/divVerent/aaaaxy/internal/fun"
	"github.com/divVerent/aaaaxy/internal/locale"
	"github.com/divVerent/aaaaxy/internal/log"
	"github.com/divVerent/aaaaxy/internal/vfs"
)

var (
	language      = flag.String("language", "auto", "language to translate the game into; if set to 'auto', it will be detected using the system locale; set to '' to not translate")
	dumpLanguages = flag.Bool("dump_languages", false, "just print the list of languages and exit")
)

func initLinguas() error {
	data, err := vfs.Load("locales", "LINGUAS")
	if err != nil {
		if *dumpLanguages {
			log.Fatalf("could not open LINGUAS file: %v", err)
		}
		log.Errorf("could not open LINGUAS file: %v", err)
		return nil
	}
	defer data.Close()
	buf, err := io.ReadAll(data)
	if err != nil {
		log.Errorf("could not read LINGUAS file: %v", err)
	}
	for _, line := range bytes.Split(buf, []byte{'\n'}) {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		l := locale.Lingua(line)
		locale.Linguas[l] = struct{}{}
		for _, member := range l.GroupMembers() {
			locale.Linguas[member] = struct{}{}
		}
	}
	// Try detecting language packs.
	domains := []string{"game"}
	levels, err := vfs.ReadDir("maps")
	if err != nil {
		log.Errorf("could not list levels: %v", err)
	}
	for _, level := range levels {
		name, isTMX := strings.CutSuffix(level, ".tmx")
		if !isTMX {
			continue
		}
		domains = append(domains, name)
	}
	sort.Strings(domains)
	log.Infof("supported locale domains: %v", domains)
	for _, domain := range domains {
		data, err = vfs.OSOpen(vfs.ExeDir, fmt.Sprintf("%s.po", domain))
		if err != nil {
			// Missing language pack is OK.
			continue
		}
		data.Close()
		log.Infof("detected an user provided language pack")
		locale.Linguas[locale.UserProvided] = struct{}{}
		break
	}
	if *dumpLanguages {
		for _, l := range locale.LinguasSorted() {
			fmt.Println(string(l))
		}
		return exitstatus.ErrRegularTermination
	}
	return nil
}

func initLocaleDomain(lang locale.Lingua, l locale.Type, domain string) {
	if lang == locale.Builtin {
		return
	}
	var data io.ReadCloser
	var err error
	if lang == locale.UserProvided {
		data, err = vfs.OSOpen(vfs.ExeDir, fmt.Sprintf("%s.po", domain))
	} else {
		data, err = vfs.Load(fmt.Sprintf("locales/%s", lang.Directory()), fmt.Sprintf("%s.po", domain))
	}
	if err != nil {
		log.Errorf("could not open %s translation for language %s: %v", domain, lang.Name(), err)
		return
	}
	defer data.Close()
	buf, err := io.ReadAll(data)
	if err != nil {
		log.Errorf("could not read %s translation for language %s: %v", domain, lang.Name(), err)
		return
	}
	l.Parse(buf)
	log.Infof("%s translated to language %s", domain, lang.Name())
}

func Init(levelName string) error {
	err := initLinguas()
	if err != nil {
		return err
	}
	locale.InitCurrent()
	_, err = forceSetLanguage(levelName, locale.Lingua(*language))
	return err
}

func SetLanguage(levelName string, lang locale.Lingua) (bool, error) {
	if lang == "auto" {
		lang = locale.Current
	}
	if locale.Active == lang && levelName == locale.LName {
		return false, nil
	}
	return forceSetLanguage(levelName, lang)
}

func forceSetLanguage(levelName string, lang locale.Lingua) (bool, error) {
	if lang == "auto" {
		lang = locale.Current
	}
	locale.ResetLanguage()
	initLocaleDomain(lang, locale.G, "game")
	initLocaleDomain(lang, locale.L, levelName)
	locale.Active = lang
	locale.LName = levelName

	err := font.SetFont(locale.ActiveFont())
	if err != nil {
		return false, err
	}

	// Now perform all replacements in locale.G.
	// In locale.L they're applied at runtime as more stuff may need filling in.
	// This must be done after setting it active, and before auditing.
	for _, t := range locale.G.GetDomain().GetTranslations() {
		if len(t.Trs) != 1 {
			// Sorry, not supporting plurals yet.
			continue
		}
		if strings.Contains(t.ID, "{{") {
			// If the LHS is a template already, no need to replace.
			continue
		}
		msgstr, ok := t.Trs[0]
		if !ok {
			// Odd - no singular form?
			continue
		}
		replacement, err := fun.TryFormatText(nil, msgstr)
		if err != nil {
			// Failed to format? This usually means a syntax error.
			// Format strings that fail due to no player state should not get here.
			// They should have been caught by the {{ check above.
			log.Fatalf("invalid msgstr: %q: %v", msgstr, err)
		}
		if replacement == msgstr {
			// No change.
			continue
		}
		locale.G.Set(t.ID, replacement)
	}
	return true, locale.Audit()
}
