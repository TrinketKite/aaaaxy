// Copyright 2021 Google LLC
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

package palette

// data contains all palettes by name.
var data = map[string]*Palette{
	// Monochrome.
	"mono": newPalette([]int{0, 15}, []uint32{
		0x000000,
		0xFFFFFF,
	}),

	// The original IBM CGA palette.
	"cga40l": newPalette([]int{0, 2, 4, 6}, []uint32{
		0x000000,
		0x00AA00,
		0xAA0000,
		0xAA5500,
	}),

	// The original IBM CGA palette at high intensity.
	"cga40h": newPalette([]int{0, 10, 12, 14}, []uint32{
		0x000000,
		0x55FF55,
		0xFF5555,
		0xFFFF55,
	}),

	// The original IBM CGA palette on NTSC.
	// curl https://upload.wikimedia.org/wikipedia/commons/7/7c/CGA_CompVsRGB_320p0.png | convert PNG:- -crop 100x180+500+10 -compress none PNM:- | tail -n +4 | uniq | awk '{ printf "0x%02X%02X%02X,\n", $1, $2, $3; }'
	// The EGA mapping here is very approximate as mapping more colors tends to look better even if not very close.
	"cga40n": newPalette([]int{0, 3, 9, 1, 6, 2, 10, 4, 7, 13, 14, 12}, []uint32{
		0x000000, // Black.
		0x0071D1, // Cyan.
		0x0071F1, // Light blue.
		0x0019AC, // Blue.
		0x954F00, // Brown.
		0x6DD441, // Green.
		0x74D461, // Light green.
		0xB82100, // Red.
		0x90A69C, // Light grey.
		0xC54E76, // Light magenta.
		0xD2ED46, // Yellow.
		0xF36800, // Light red.
		0x97A6BB, // Other light grey. Unmapped.
		0xA27B1C, // Light brown. Unmapped.
		0xCBED26, // Other yellow. Unmapped.
		0xFF9501, // Orange. Unmapped.
	}),

	// The alternate IBM CGA palette.
	"cga41l": newPalette([]int{0, 3, 5, 7}, []uint32{
		0x000000,
		0x00AAAA,
		0xAA00AA,
		0xAAAAAA,
	}),

	// The alternate IBM CGA palette at high intensity.
	"cga41h": newPalette([]int{0, 11, 13, 15}, []uint32{
		0x000000,
		0x55FFFF,
		0xFF55FF,
		0xFFFFFF,
	}),

	// The alternate IBM CGA palette on NTSC.
	// curl https://upload.wikimedia.org/wikipedia/commons/c/c5/CGA_CompVsRGB_320p1.png | convert PNG:- -crop 100x180+500+10 -compress none PNM:- | tail -n +4 | uniq | awk '{ printf "0x%02X%02X%02X,\n", $1, $2, $3; }'
	// The EGA mapping here is very approximate as mapping more colors tends to look better even if not very close.
	"cga41n": newPalette([]int{0, 3, 1, 6, 9, 7, 11, 4, 13, 10, 12, 15}, []uint32{
		0x000000, // Black.
		0x009AFF, // Cyan.
		0x0042FF, // Blue.
		0xAA4C00, // Brown.
		0xA7CDFF, // Light blue.
		0xB9A2AD, // Light grey.
		0x96F0FF, // Light cyan.
		0xCD1F00, // Red.
		0xDC75FF, // Light magenta.
		0xEDFFCC, // Light green.
		0xFFB2A6, // Light red.
		0xFFFFFF, // White.
		0x0090FF, // Other cyan. Unmapped.
		0x84FAD2, // Other light cyan. Unmapped.
		0xB9C3FF, // Other light grey. Unmapped.
		0xFF5C00, // Orange. Unmapped.
	}),

	// The "monochrome" IBM CGA palette.
	"cga5l": newPalette([]int{0, 3, 4, 7}, []uint32{
		0x000000,
		0x00AAAA,
		0xAA0000,
		0xAAAAAA,
	}),

	// The "monochrome" IBM CGA palette at high intensity.
	"cga5h": newPalette([]int{0, 11, 12, 15}, []uint32{
		0x000000,
		0x55FFFF,
		0xFF5555,
		0xFFFFFF,
	}),

	// The palette one gets when using the CGA monochrome mode on NTSC while forcing the colorburst signal.
	// curl https://upload.wikimedia.org/wikipedia/commons/f/fb/CGA_CompVsRGB_640.png | convert PNG:- -crop 100x360+1000+20 -compress none PNM:- | tail -n +4 | uniq | awk '{ printf "0x%02X%02X%02X,\n", $1, $2, $3; }'
	// The EGA mapping here is very approximate as mapping more colors tends to look better even if not very close.
	"cga6n": newPalette([]int{0, 2, 9, 10, 1, 11, 7, 4, 14, 13, 12, 15}, []uint32{
		0x000000, // Black.
		0x006E31, // Green.
		0x008AFF, // Light blue.
		0x00DB00, // Light green.
		0x3109FF, // Blue.
		0x45F7BB, // Light cyan.
		0x767676, // Light grey.
		0xA70031, // Red.
		0xBBE400, // Yellow.
		0xEC11FF, // Light magenta.
		0xEC6300, // Light red.
		0xFFFFFF, // White.
		0x315A00, // Other green. Unmapped.
		0xBB92FF, // Pink. Unmapped.
		0xFF7FBB, // Light pink. Unmapped.
		0x767676, // Other light grey. Redundant.
	}),

	// The original IBM EGA palette.
	"ega": newPalette([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, egaColors[:]),

	// EGA but only the colors 0 to 7.
	"egalow": newPalette([]int{0, 1, 2, 3, 4, 5, 6, 7}, []uint32{
		0x000000,
		0x0000AA,
		0x00AA00,
		0x00AAAA,
		0xAA0000,
		0xAA00AA,
		0xAA5500,
		0xAAAAAA,
	}),

	// EGA but only the grey tones.
	"egamono": newPalette([]int{0, 7, 8, 15}, []uint32{
		0x000000,
		0xAAAAAA,
		0x555555,
		0xFFFFFF,
	}),

	// XTerm's ANSI palette.
	"xterm": newPalette([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []uint32{
		0x000000,
		0x000080,
		0x008000,
		0x008080,
		0x800000,
		0x800080,
		0x808000,
		0xC0C0C0,
		0x808080,
		0x0000FF,
		0x00FF00,
		0x00FFFF,
		0xFF0000,
		0xFF00FF,
		0xFFFF00,
		0xFFFFFF,
	}),

	// My favorite ANSI palette variant. Good color contrast.
	// General rule: greys like EGA, but other colors are all
	// of form:
	// - one FF two 00 (color cube corner)
	// - two 80 one 00 (color cube side midpoint)
	// - two 80 one FF (color cube side midpoint)
	// - one 00 two FF (color cube corner)
	"div0": newPalette([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []uint32{
		0x000000,
		0x0000FF,
		0x00FF00,
		0x008080,
		0xFF0000,
		0x800080,
		0x808000,
		0xAAAAAA,
		0x555555,
		0x8080FF,
		0x80FF80,
		0x00FFFF,
		0xFF8080,
		0xFF00FF,
		0xFFFF00,
		0xFFFFFF,
	}),

	// Atari ST palette.
	// Has 7 levels in each component, representing as 00 24 49 6D 92 B6 DB FF.
	"atarist": newPalette([]int{15, 0, 12, 10, 9, 1, 4, 2, 7, 8, 11, 3, 13, 5, 14, 6}, []uint32{
		0xFFFFFF,
		0x000000,
		0xFF0000,
		0x00FF00,
		0x0000FF,
		0x000092,
		0x924900,
		0x009200,
		0xB6B6B6,
		0x494949,
		0x00FFFF,
		0x009292,
		0xFF00FF,
		0x920092,
		0xFFFF00,
		0x929200,
	}),

	// Macintosh II palette. Lacks dark tones of blue and cyan.
	"macii": newPalette([]int{15, 14, 12, 13, 5, 9, 11, 10, 2, 4, 6, 7, 8, 0}, []uint32{
		0xFFFFFF, // White.
		0xFBF305, // Yellow.
		0xDD0907, // Light red.
		0xF20884, // Light magenta.
		0x4700A5, // Purple. Using as magenta.
		0x0000D3, // Light blue.
		0x02ABEA, // Light cyan.
		0x1FB714, // Light green.
		0x006412, // Green.
		0x562C05, // Brown. Using as red.
		0x90713A, // Tan. Using as brown.
		0xC0C0C0, // Light grey.
		0x404040, // Dark grey.
		0x000000, // Black.
		0xFF6403, // Orange. Unmapped.
		0x808080, // Grey. Unmapped.
	}),

	// The original IBM VGA palette, with colors too close to EGA colors commented out.
	"vga": newPalette([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []uint32{
		0x000000,
		0x0000AA,
		0x00AA00,
		0x00AAAA,
		0xAA0000,
		0xAA00AA,
		0xAA5500,
		0xAAAAAA,
		0x555555,
		0x5555FF,
		0x55FF55,
		0x55FFFF,
		0xFF5555,
		0xFF55FF,
		0xFFFF55,
		0xFFFFFF,
		0x000000,
		0x141414,
		0x202020,
		0x2C2C2C,
		0x383838,
		0x444444,
		0x505050,
		0x616161,
		0x717171,
		0x818181,
		0x919191,
		0xA1A1A1,
		0xB6B6B6,
		0xCACACA,
		0xE2E2E2,
		0xFFFFFF,
		0x0000FF,
		0x4000FF,
		0x7D00FF,
		0xBE00FF,
		0xFF00FF,
		0xFF00BE,
		0xFF007D,
		0xFF0040,
		0xFF0000,
		0xFF4000,
		0xFF7D00,
		0xFFBE00,
		0xFFFF00,
		0xBEFF00,
		0x7DFF00,
		0x40FF00,
		0x00FF00,
		0x00FF40,
		0x00FF7D,
		0x00FFBE,
		0x00FFFF,
		0x00BEFF,
		0x007DFF,
		0x0040FF,
		0x7D7DFF,
		0x9D7DFF,
		0xBE7DFF,
		0xDE7DFF,
		0xFF7DFF,
		0xFF7DDE,
		0xFF7DBE,
		0xFF7D9D,
		0xFF7D7D,
		0xFF9D7D,
		0xFFBE7D,
		0xFFDE7D,
		0xFFFF7D,
		0xDEFF7D,
		0xBEFF7D,
		0x9DFF7D,
		0x7DFF7D,
		0x7DFF9D,
		0x7DFFBE,
		0x7DFFDE,
		0x7DFFFF,
		0x7DDEFF,
		0x7DBEFF,
		0x7D9DFF,
		0xB6B6FF,
		0xC6B6FF,
		0xDAB6FF,
		0xEAB6FF,
		0xFFB6FF,
		0xFFB6EA,
		0xFFB6DA,
		0xFFB6C6,
		0xFFB6B6,
		0xFFC6B6,
		0xFFDAB6,
		0xFFEAB6,
		0xFFFFB6,
		0xEAFFB6,
		0xDAFFB6,
		0xC6FFB6,
		0xB6FFB6,
		0xB6FFC6,
		0xB6FFDA,
		0xB6FFEA,
		0xB6FFFF,
		0xB6EAFF,
		0xB6DAFF,
		0xB6C6FF,
		0x000071,
		0x1C0071,
		0x380071,
		0x550071,
		0x710071,
		0x710055,
		0x710038,
		0x71001C,
		0x710000,
		0x711C00,
		0x713800,
		0x715500,
		0x717100,
		0x557100,
		0x387100,
		0x1C7100,
		0x007100,
		0x00711C,
		0x007138,
		0x007155,
		0x007171,
		0x005571,
		0x003871,
		0x001C71,
		0x383871,
		0x443871,
		0x553871,
		0x613871,
		0x713871,
		0x713861,
		0x713855,
		0x713844,
		0x713838,
		0x714438,
		0x715538,
		0x716138,
		0x717138,
		0x617138,
		0x557138,
		0x447138,
		0x387138,
		0x387144,
		0x387155,
		0x387161,
		0x387171,
		0x386171,
		0x385571,
		0x384471,
		0x505071,
		0x595071,
		0x615071,
		0x695071,
		0x715071,
		0x715069,
		0x715061,
		0x715059,
		0x715050,
		0x715950,
		0x716150,
		0x716950,
		0x717150,
		0x697150,
		0x617150,
		0x597150,
		0x507150,
		0x507159,
		0x507161,
		0x507169,
		0x507171,
		0x506971,
		0x506171,
		0x505971,
		0x000040,
		0x100040,
		0x200040,
		0x300040,
		0x400040,
		0x400030,
		0x400020,
		0x400010,
		0x400000,
		0x401000,
		0x402000,
		0x403000,
		0x404000,
		0x304000,
		0x204000,
		0x104000,
		0x004000,
		0x004010,
		0x004020,
		0x004030,
		0x004040,
		0x003040,
		0x002040,
		0x001040,
		0x202040,
		0x282040,
		0x302040,
		0x382040,
		0x402040,
		0x402038,
		0x402030,
		0x402028,
		0x402020,
		0x402820,
		0x403020,
		0x403820,
		0x404020,
		0x384020,
		0x304020,
		0x284020,
		0x204020,
		0x204028,
		0x204030,
		0x204038,
		0x204040,
		0x203840,
		0x203040,
		0x202840,
		0x2C2C40,
		0x302C40,
		0x342C40,
		0x3C2C40,
		0x402C40,
		0x402C3C,
		0x402C34,
		0x402C30,
		0x402C2C,
		0x40302C,
		0x40342C,
		0x403C2C,
		0x40402C,
		0x3C402C,
		0x34402C,
		0x30402C,
		0x2C402C,
		0x2C4030,
		0x2C4034,
		0x2C403C,
		0x2C4040,
		0x2C3C40,
		0x2C3440,
		0x2C3040,
		0x000000,
		0x000000,
		0x000000,
		0x000000,
		0x000000,
		0x000000,
		0x000000,
		0x000000,
	}),

	// Quake's palette. Has been put in the public domain by John Carmack.
	"quake": newPalette([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []uint32{
		0x000000, // Black.
		0x2B2BAF, // Blue.
		0x475B4F, // Green.
		0x63638B, // Cyan.
		0x570000, // Red.
		0x7F4B5F, // Magenta.
		0x634B1F, // Brown.
		0x9B9B9B, // Light grey.
		0x4B4B4B, // Dark grey.
		0x0000FF, // Light blue.
		0x6F837B, // Light green.
		0x8B8BCB, // Light cyan.
		0x7F0000, // Light red.
		0xBB739F, // Light magenta.
		0xFFF31B, // Yellow.
		0xFFFFFF, // White.

		0x0F0F0F,
		0x1F1F1F,
		0x2F2F2F,
		0x3F3F3F,
		0x5B5B5B,
		0x6B6B6B,
		0x7B7B7B,
		0x8B8B8B,
		0xABABAB,
		0xBBBBBB,
		0xCBCBCB,
		0xDBDBDB,
		0xEBEBEB,
		0x0F0B07,
		0x170F0B,
		0x1F170B,
		0x271B0F,
		0x2F2313,
		0x372B17,
		0x3F2F17,
		0x4B371B,
		0x533B1B,
		0x5B431F,
		0x6B531F,
		0x73571F,
		0x7B5F23,
		0x836723,
		0x8F6F23,
		0x0B0B0F,
		0x13131B,
		0x1B1B27,
		0x272733,
		0x2F2F3F,
		0x37374B,
		0x3F3F57,
		0x474767,
		0x4F4F73,
		0x5B5B7F,
		0x6B6B97,
		0x7373A3,
		0x7B7BAF,
		0x8383BB,
		0x000000,
		0x070700,
		0x0B0B00,
		0x131300,
		0x1B1B00,
		0x232300,
		0x2B2B07,
		0x2F2F07,
		0x373707,
		0x3F3F07,
		0x474707,
		0x4B4B0B,
		0x53530B,
		0x5B5B0B,
		0x63630B,
		0x6B6B0F,
		0x070000,
		0x0F0000,
		0x170000,
		0x1F0000,
		0x270000,
		0x2F0000,
		0x370000,
		0x3F0000,
		0x470000,
		0x4F0000,
		0x5F0000,
		0x670000,
		0x6F0000,
		0x770000,
		0x131300,
		0x1B1B00,
		0x232300,
		0x2F2B00,
		0x372F00,
		0x433700,
		0x4B3B07,
		0x574307,
		0x5F4707,
		0x6B4B0B,
		0x77530F,
		0x835713,
		0x8B5B13,
		0x975F1B,
		0xA3631F,
		0xAF6723,
		0x231307,
		0x2F170B,
		0x3B1F0F,
		0x4B2313,
		0x572B17,
		0x632F1F,
		0x733723,
		0x7F3B2B,
		0x8F4333,
		0x9F4F33,
		0xAF632F,
		0xBF772F,
		0xCF8F2B,
		0xDFAB27,
		0xEFCB1F,
		0x0B0700,
		0x1B1300,
		0x2B230F,
		0x372B13,
		0x47331B,
		0x533723,
		0x633F2B,
		0x6F4733,
		0x7F533F,
		0x8B5F47,
		0x9B6B53,
		0xA77B5F,
		0xB7876B,
		0xC3937B,
		0xD3A38B,
		0xE3B397,
		0xAB8BA3,
		0x9F7F97,
		0x937387,
		0x8B677B,
		0x7F5B6F,
		0x775363,
		0x6B4B57,
		0x5F3F4B,
		0x573743,
		0x4B2F37,
		0x43272F,
		0x371F23,
		0x2B171B,
		0x231313,
		0x170B0B,
		0x0F0707,
		0xAF6B8F,
		0xA35F83,
		0x975777,
		0x8B4F6B,
		0x734353,
		0x6B3B4B,
		0x5F333F,
		0x532B37,
		0x47232B,
		0x3B1F23,
		0x2F171B,
		0x231313,
		0x170B0B,
		0x0F0707,
		0xDBC3BB,
		0xCBB3A7,
		0xBFA39B,
		0xAF978B,
		0xA3877B,
		0x977B6F,
		0x876F5F,
		0x7B6353,
		0x6B5747,
		0x5F4B3B,
		0x533F33,
		0x433327,
		0x372B1F,
		0x271F17,
		0x1B130F,
		0x0F0B07,
		0x677B6F,
		0x5F7367,
		0x576B5F,
		0x4F6357,
		0x3F5347,
		0x374B3F,
		0x2F4337,
		0x2B3B2F,
		0x233327,
		0x1F2B1F,
		0x172317,
		0x0F1B13,
		0x0B130B,
		0x070B07,
		0xFFF31B,
		0xEFDF17,
		0xDBCB13,
		0xCBB70F,
		0xBBA70F,
		0xAB970B,
		0x9B8307,
		0x8B7307,
		0x7B6307,
		0x6B5300,
		0x5B4700,
		0x4B3700,
		0x3B2B00,
		0x2B1F00,
		0x1B0F00,
		0x0B0700,
		0x0B0BEF,
		0x1313DF,
		0x1B1BCF,
		0x2323BF,
		0x2F2F9F,
		0x2F2F8F,
		0x2F2F6F,
		0x2F2F5F,
		0x2B2B4F,
		0x23233F,
		0x1B1B2F,
		0x13131F,
		0x0B0B0F,
		0x2B0000,
		0x3B0000,
		0x4B0700,
		0x5F0700,
		0x6F0F00,
		0x7F1707,
		0x931F07,
		0xA3270B,
		0xB7330F,
		0xC34B1B,
		0xCF632B,
		0xDB7F3B,
		0xE3974F,
		0xE7AB5F,
		0xEFBF77,
		0xF7D38B,
		0xA77B3B,
		0xB79B37,
		0xC7C337,
		0xE7E357,
		0x7FBFFF,
		0xABE7FF,
		0xD7FFFF,
		0x670000,
		0x8B0000,
		0xB30000,
		0xD70000,
		0xFF0000,
		0xFFF393,
		0xFFF7C7,
		0x9F5B53,
	}),

	// A wellknown subset of the NES palette.
	// In fact, the exact set of colors visible while 1-1 is loaded.
	// Only colors not mapped: blue, cyan, dark grey, light grey.
	"smb": newPalette([]int{9, 10, 5, 2, 14, 0, 6, 13, 12, 15, 11, 4}, []uint32{
		nesColors[0x22], // Light blue.
		nesColors[0x29], // Light green.
		nesColors[0x16], // Magenta
		nesColors[0x1A], // Green.
		nesColors[0x27], // Yellow.
		nesColors[0x0F], // Black.
		nesColors[0x18], // Brown.
		nesColors[0x36], // Light magenta.
		nesColors[0x17], // Light red.
		nesColors[0x30], // White.
		nesColors[0x21], // Light cyan.
		nesColors[0x07], // Red.
	}),

	// Same as "smb" but with all missing colors added.
	"nes": newPalette([]int{9, 10, 5, 2, 14, 0, 6, 13, 12, 15, 11, 4, 1, 3, 7, 8}, append([]uint32{
		nesColors[0x22], // Light blue.
		nesColors[0x29], // Light green.
		nesColors[0x16], // Magenta
		nesColors[0x1A], // Green.
		nesColors[0x27], // Yellow.
		nesColors[0x0F], // Black.
		nesColors[0x18], // Brown.
		nesColors[0x36], // Light magenta.
		nesColors[0x17], // Light red.
		nesColors[0x30], // White.
		nesColors[0x21], // Light cyan.
		nesColors[0x07], // Red.

		// Other colors to fill up the ANSI palette.
		nesColors[0x02], // Blue.
		nesColors[0x11], // Cyan.
		nesColors[0x10], // Light grey.
		nesColors[0x00], // Dark grey.
	}, nesColors[:]...)),

	// Gameboy?
	"gb": newPalette([]int{0, 8, 7, 15}, []uint32{
		0x081820,
		0x346856,
		0x88C070,
		0xE0F8D0,
	}),

	// C64 palette.
	// NOT protecting "orange" and one of the greys (listed at the end).
	// Lacking equivalents of EGA bright cyan and bright pink.
	// However, as this game uses the bright colors more than the dark colors,
	// decided to map the bright colors to the C64 colors, noted below.
	"c64": newPalette([]int{0, 15, 4, 11, 13, 2, 1, 14, 6, 12, 8, 7, 10, 9}, []uint32{
		0x000000, // Black.
		0xFFFFFF, // White.
		0x883932, // Normal red.
		0x67B6BD, // Normal cyan mapped as bright.
		0x8B3F96, // Normal pink mapped as bright.
		0x55A049, // Normal green.
		0x40318D, // Normal blue.
		0xBFCE72, // Yellow.
		0x574200, // Normal brown.
		0xB86962, // Bright red.
		0x505050, // Dark grey.
		0x9F9F9F, // Bright grey.
		0x94E089, // Bright green.
		0x7869C4, // Bright blue.
		0x787878, // 50% grey. Unmapped.
		0x8B5429, // Orange. Unmapped.
	}),

	// Web safe 216 colors palette, actually a 6x6x6 color cube.
	// Dither everywhere.
	"web": newPalette(nil, []uint32{
		0x000000,
		0x000033,
		0x000066,
		0x000099,
		0x0000CC,
		0x0000FF,
		0x003300,
		0x003333,
		0x003366,
		0x003399,
		0x0033CC,
		0x0033FF,
		0x006600,
		0x006633,
		0x006666,
		0x006699,
		0x0066CC,
		0x0066FF,
		0x009900,
		0x009933,
		0x009966,
		0x009999,
		0x0099CC,
		0x0099FF,
		0x00CC00,
		0x00CC33,
		0x00CC66,
		0x00CC99,
		0x00CCCC,
		0x00CCFF,
		0x00FF00,
		0x00FF33,
		0x00FF66,
		0x00FF99,
		0x00FFCC,
		0x00FFFF,
		0x330000,
		0x330033,
		0x330066,
		0x330099,
		0x3300CC,
		0x3300FF,
		0x333300,
		0x333333,
		0x333366,
		0x333399,
		0x3333CC,
		0x3333FF,
		0x336600,
		0x336633,
		0x336666,
		0x336699,
		0x3366CC,
		0x3366FF,
		0x339900,
		0x339933,
		0x339966,
		0x339999,
		0x3399CC,
		0x3399FF,
		0x33CC00,
		0x33CC33,
		0x33CC66,
		0x33CC99,
		0x33CCCC,
		0x33CCFF,
		0x33FF00,
		0x33FF33,
		0x33FF66,
		0x33FF99,
		0x33FFCC,
		0x33FFFF,
		0x660000,
		0x660033,
		0x660066,
		0x660099,
		0x6600CC,
		0x6600FF,
		0x663300,
		0x663333,
		0x663366,
		0x663399,
		0x6633CC,
		0x6633FF,
		0x666600,
		0x666633,
		0x666666,
		0x666699,
		0x6666CC,
		0x6666FF,
		0x669900,
		0x669933,
		0x669966,
		0x669999,
		0x6699CC,
		0x6699FF,
		0x66CC00,
		0x66CC33,
		0x66CC66,
		0x66CC99,
		0x66CCCC,
		0x66CCFF,
		0x66FF00,
		0x66FF33,
		0x66FF66,
		0x66FF99,
		0x66FFCC,
		0x66FFFF,
		0x990000,
		0x990033,
		0x990066,
		0x990099,
		0x9900CC,
		0x9900FF,
		0x993300,
		0x993333,
		0x993366,
		0x993399,
		0x9933CC,
		0x9933FF,
		0x996600,
		0x996633,
		0x996666,
		0x996699,
		0x9966CC,
		0x9966FF,
		0x999900,
		0x999933,
		0x999966,
		0x999999,
		0x9999CC,
		0x9999FF,
		0x99CC00,
		0x99CC33,
		0x99CC66,
		0x99CC99,
		0x99CCCC,
		0x99CCFF,
		0x99FF00,
		0x99FF33,
		0x99FF66,
		0x99FF99,
		0x99FFCC,
		0x99FFFF,
		0xCC0000,
		0xCC0033,
		0xCC0066,
		0xCC0099,
		0xCC00CC,
		0xCC00FF,
		0xCC3300,
		0xCC3333,
		0xCC3366,
		0xCC3399,
		0xCC33CC,
		0xCC33FF,
		0xCC6600,
		0xCC6633,
		0xCC6666,
		0xCC6699,
		0xCC66CC,
		0xCC66FF,
		0xCC9900,
		0xCC9933,
		0xCC9966,
		0xCC9999,
		0xCC99CC,
		0xCC99FF,
		0xCCCC00,
		0xCCCC33,
		0xCCCC66,
		0xCCCC99,
		0xCCCCCC,
		0xCCCCFF,
		0xCCFF00,
		0xCCFF33,
		0xCCFF66,
		0xCCFF99,
		0xCCFFCC,
		0xCCFFFF,
		0xFF0000,
		0xFF0033,
		0xFF0066,
		0xFF0099,
		0xFF00CC,
		0xFF00FF,
		0xFF3300,
		0xFF3333,
		0xFF3366,
		0xFF3399,
		0xFF33CC,
		0xFF33FF,
		0xFF6600,
		0xFF6633,
		0xFF6666,
		0xFF6699,
		0xFF66CC,
		0xFF66FF,
		0xFF9900,
		0xFF9933,
		0xFF9966,
		0xFF9999,
		0xFF99CC,
		0xFF99FF,
		0xFFCC00,
		0xFFCC33,
		0xFFCC66,
		0xFFCC99,
		0xFFCCCC,
		0xFFCCFF,
		0xFFFF00,
		0xFFFF33,
		0xFFFF66,
		0xFFFF99,
		0xFFFFCC,
		0xFFFFFF,
	}),

	// 2x2x2 color cube. Just eight pure colors.
	"2x2x2": newPalette([]int{0, 9, 10, 11, 12, 13, 14, 15}, []uint32{
		0x000000,
		0x0000FF,
		0x00FF00,
		0x00FFFF,
		0xFF0000,
		0xFF00FF,
		0xFFFF00,
		0xFFFFFF,
	}),

	// 4x4x4 color cube. For those who just like the dither.
	// Dither everywhere.
	"4x4x4": newPalette(nil, []uint32{
		0x000000,
		0x000055,
		0x0000AA,
		0x0000FF,
		0x005500,
		0x005555,
		0x0055AA,
		0x0055FF,
		0x00AA00,
		0x00AA55,
		0x00AAAA,
		0x00AAFF,
		0x00FF00,
		0x00FF55,
		0x00FFAA,
		0x00FFFF,
		0x550000,
		0x550055,
		0x5500AA,
		0x5500FF,
		0x555500,
		0x555555,
		0x5555AA,
		0x5555FF,
		0x55AA00,
		0x55AA55,
		0x55AAAA,
		0x55AAFF,
		0x55FF00,
		0x55FF55,
		0x55FFAA,
		0x55FFFF,
		0xAA0000,
		0xAA0055,
		0xAA00AA,
		0xAA00FF,
		0xAA5500,
		0xAA5555,
		0xAA55AA,
		0xAA55FF,
		0xAAAA00,
		0xAAAA55,
		0xAAAAAA,
		0xAAAAFF,
		0xAAFF00,
		0xAAFF55,
		0xAAFFAA,
		0xAAFFFF,
		0xFF0000,
		0xFF0055,
		0xFF00AA,
		0xFF00FF,
		0xFF5500,
		0xFF5555,
		0xFF55AA,
		0xFF55FF,
		0xFFAA00,
		0xFFAA55,
		0xFFAAAA,
		0xFFAAFF,
		0xFFFF00,
		0xFFFF55,
		0xFFFFAA,
		0xFFFFFF,
	}),

	// 7x7x4 color "cube". Cleanest colors at 256c. Doesn't do blue gradients.
	// Dither everywhere.
	"7x7x4": newPalette(nil, []uint32{
		0x000000,
		0x000055,
		0x0000AA,
		0x0000FF,
		0x002A00,
		0x002A55,
		0x002AAA,
		0x002AFF,
		0x005500,
		0x005555,
		0x0055AA,
		0x0055FF,
		0x007F00,
		0x007F55,
		0x007FAA,
		0x007FFF,
		0x00AA00,
		0x00AA55,
		0x00AAAA,
		0x00AAFF,
		0x00D400,
		0x00D455,
		0x00D4AA,
		0x00D4FF,
		0x00FF00,
		0x00FF55,
		0x00FFAA,
		0x00FFFF,
		0x2A0000,
		0x2A0055,
		0x2A00AA,
		0x2A00FF,
		0x2A2A00,
		0x2A2A55,
		0x2A2AAA,
		0x2A2AFF,
		0x2A5500,
		0x2A5555,
		0x2A55AA,
		0x2A55FF,
		0x2A7F00,
		0x2A7F55,
		0x2A7FAA,
		0x2A7FFF,
		0x2AAA00,
		0x2AAA55,
		0x2AAAAA,
		0x2AAAFF,
		0x2AD400,
		0x2AD455,
		0x2AD4AA,
		0x2AD4FF,
		0x2AFF00,
		0x2AFF55,
		0x2AFFAA,
		0x2AFFFF,
		0x550000,
		0x550055,
		0x5500AA,
		0x5500FF,
		0x552A00,
		0x552A55,
		0x552AAA,
		0x552AFF,
		0x555500,
		0x555555,
		0x5555AA,
		0x5555FF,
		0x557F00,
		0x557F55,
		0x557FAA,
		0x557FFF,
		0x55AA00,
		0x55AA55,
		0x55AAAA,
		0x55AAFF,
		0x55D400,
		0x55D455,
		0x55D4AA,
		0x55D4FF,
		0x55FF00,
		0x55FF55,
		0x55FFAA,
		0x55FFFF,
		0x7F0000,
		0x7F0055,
		0x7F00AA,
		0x7F00FF,
		0x7F2A00,
		0x7F2A55,
		0x7F2AAA,
		0x7F2AFF,
		0x7F5500,
		0x7F5555,
		0x7F55AA,
		0x7F55FF,
		0x7F7F00,
		0x7F7F55,
		0x7F7FAA,
		0x7F7FFF,
		0x7FAA00,
		0x7FAA55,
		0x7FAAAA,
		0x7FAAFF,
		0x7FD400,
		0x7FD455,
		0x7FD4AA,
		0x7FD4FF,
		0x7FFF00,
		0x7FFF55,
		0x7FFFAA,
		0x7FFFFF,
		0xAA0000,
		0xAA0055,
		0xAA00AA,
		0xAA00FF,
		0xAA2A00,
		0xAA2A55,
		0xAA2AAA,
		0xAA2AFF,
		0xAA5500,
		0xAA5555,
		0xAA55AA,
		0xAA55FF,
		0xAA7F00,
		0xAA7F55,
		0xAA7FAA,
		0xAA7FFF,
		0xAAAA00,
		0xAAAA55,
		0xAAAAAA,
		0xAAAAFF,
		0xAAD400,
		0xAAD455,
		0xAAD4AA,
		0xAAD4FF,
		0xAAFF00,
		0xAAFF55,
		0xAAFFAA,
		0xAAFFFF,
		0xD40000,
		0xD40055,
		0xD400AA,
		0xD400FF,
		0xD42A00,
		0xD42A55,
		0xD42AAA,
		0xD42AFF,
		0xD45500,
		0xD45555,
		0xD455AA,
		0xD455FF,
		0xD47F00,
		0xD47F55,
		0xD47FAA,
		0xD47FFF,
		0xD4AA00,
		0xD4AA55,
		0xD4AAAA,
		0xD4AAFF,
		0xD4D400,
		0xD4D455,
		0xD4D4AA,
		0xD4D4FF,
		0xD4FF00,
		0xD4FF55,
		0xD4FFAA,
		0xD4FFFF,
		0xFF0000,
		0xFF0055,
		0xFF00AA,
		0xFF00FF,
		0xFF2A00,
		0xFF2A55,
		0xFF2AAA,
		0xFF2AFF,
		0xFF5500,
		0xFF5555,
		0xFF55AA,
		0xFF55FF,
		0xFF7F00,
		0xFF7F55,
		0xFF7FAA,
		0xFF7FFF,
		0xFFAA00,
		0xFFAA55,
		0xFFAAAA,
		0xFFAAFF,
		0xFFD400,
		0xFFD455,
		0xFFD4AA,
		0xFFD4FF,
		0xFFFF00,
		0xFFFF55,
		0xFFFFAA,
		0xFFFFFF,
	}),

	// A flag.
	"ua3": newPalette([]int{0, 9, 14}, []uint32{
		0x000000,
		0x0057B8,
		0xFFD700,
	}),

	// Another flag.
	"de3": newPalette([]int{0, 12, 14}, []uint32{
		0x000000,
		0xFF0000,
		0xFFCC00,
	}),

	// Another flag.
	"us4": newPalette([]int{0, 15, 12, 9}, []uint32{
		0x000000,
		0xFFFFFF,
		0xB22234,
		0x3C3B6E,
	}),

	// Another flag.
	"ru4": newPalette([]int{0, 15, 1, 12}, []uint32{
		0x000000,
		0xFFFFFF,
		0x0032A0,
		0xDA291C,
	}),
}

func init() {
	for name, p := range data {
		p.name = name
	}
}
