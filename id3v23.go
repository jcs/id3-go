// Copyright 2013 Michael Yang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package id3

var (
	V3FrameTypeMap = map[string]FrameType{
		"AENC": FrameType{id: "AENC", description: "Audio encryption", constructor: NewDataFrame},
		"APIC": FrameType{id: "APIC", description: "Attached picture", constructor: NewImageFrame},
		"COMM": FrameType{id: "COMM", description: "Comments", constructor: NewDataFrame},
		"COMR": FrameType{id: "COMR", description: "Commercial frame", constructor: NewDataFrame},
		"ENCR": FrameType{id: "ENCR", description: "Encryption method registration", constructor: NewDataFrame},
		"EQUA": FrameType{id: "EQUA", description: "Equalization", constructor: NewDataFrame},
		"ETCO": FrameType{id: "ETCO", description: "Event timing codes", constructor: NewDataFrame},
		"GEOB": FrameType{id: "GEOB", description: "General encapsulated object", constructor: NewDataFrame},
		"GRID": FrameType{id: "GRID", description: "Group identification registration", constructor: NewDataFrame},
		"IPLS": FrameType{id: "IPLS", description: "Involved people list", constructor: NewDataFrame},
		"LINK": FrameType{id: "LINK", description: "Linked information", constructor: NewDataFrame},
		"MCDI": FrameType{id: "MCDI", description: "Music CD identifier", constructor: NewDataFrame},
		"MLLT": FrameType{id: "MLLT", description: "MPEG location lookup table", constructor: NewDataFrame},
		"OWNE": FrameType{id: "OWNE", description: "Ownership frame", constructor: NewDataFrame},
		"PRIV": FrameType{id: "PRIV", description: "Private frame", constructor: NewDataFrame},
		"PCNT": FrameType{id: "PCNT", description: "Play counter", constructor: NewDataFrame},
		"POPM": FrameType{id: "POPM", description: "Popularimeter", constructor: NewDataFrame},
		"POSS": FrameType{id: "POSS", description: "Position synchronisation frame", constructor: NewDataFrame},
		"RBUF": FrameType{id: "RBUF", description: "Recommended buffer size", constructor: NewDataFrame},
		"RVAD": FrameType{id: "RVAD", description: "Relative volume adjustment", constructor: NewDataFrame},
		"RVRB": FrameType{id: "RVRB", description: "Reverb", constructor: NewDataFrame},
		"SYLT": FrameType{id: "SYLT", description: "Synchronized lyric/text", constructor: NewDataFrame},
		"SYTC": FrameType{id: "SYTC", description: "Synchronized tempo codes", constructor: NewDataFrame},
		"TALB": FrameType{id: "TALB", description: "Album/Movie/Show title", constructor: NewTextFrame},
		"TBPM": FrameType{id: "TBPM", description: "BPM (beats per minute)", constructor: NewTextFrame},
		"TCOM": FrameType{id: "TCOM", description: "Composer", constructor: NewTextFrame},
		"TCON": FrameType{id: "TCON", description: "Content type", constructor: NewTextFrame},
		"TCOP": FrameType{id: "TCOP", description: "Copyright message", constructor: NewTextFrame},
		"TDAT": FrameType{id: "TDAT", description: "Date", constructor: NewTextFrame},
		"TDLY": FrameType{id: "TDLY", description: "Playlist delay", constructor: NewTextFrame},
		"TENC": FrameType{id: "TENC", description: "Encoded by", constructor: NewTextFrame},
		"TEXT": FrameType{id: "TEXT", description: "Lyricist/Text writer", constructor: NewTextFrame},
		"TFLT": FrameType{id: "TFLT", description: "File type", constructor: NewTextFrame},
		"TIME": FrameType{id: "TIME", description: "Time", constructor: NewTextFrame},
		"TIT1": FrameType{id: "TIT1", description: "Content group description", constructor: NewTextFrame},
		"TIT2": FrameType{id: "TIT2", description: "Title/songname/content description", constructor: NewTextFrame},
		"TIT3": FrameType{id: "TIT3", description: "Subtitle/Description refinement", constructor: NewTextFrame},
		"TKEY": FrameType{id: "TKEY", description: "Initial key", constructor: NewTextFrame},
		"TLAN": FrameType{id: "TLAN", description: "Language(s)", constructor: NewTextFrame},
		"TLEN": FrameType{id: "TLEN", description: "Length", constructor: NewTextFrame},
		"TMED": FrameType{id: "TMED", description: "Media type", constructor: NewTextFrame},
		"TOAL": FrameType{id: "TOAL", description: "Original album/movie/show title", constructor: NewTextFrame},
		"TOFN": FrameType{id: "TOFN", description: "Original filename", constructor: NewTextFrame},
		"TOLY": FrameType{id: "TOLY", description: "Original lyricist(s)/text writer(s)", constructor: NewTextFrame},
		"TOPE": FrameType{id: "TOPE", description: "Original artist(s)/performer(s)", constructor: NewTextFrame},
		"TORY": FrameType{id: "TORY", description: "Original release year", constructor: NewTextFrame},
		"TOWN": FrameType{id: "TOWN", description: "File owner/licensee", constructor: NewTextFrame},
		"TPE1": FrameType{id: "TPE1", description: "Lead performer(s)/Soloist(s)", constructor: NewTextFrame},
		"TPE2": FrameType{id: "TPE2", description: "Band/orchestra/accompaniment", constructor: NewTextFrame},
		"TPE3": FrameType{id: "TPE3", description: "Conductor/performer refinement", constructor: NewTextFrame},
		"TPE4": FrameType{id: "TPE4", description: "Interpreted, remixed, or otherwise modified by", constructor: NewTextFrame},
		"TPOS": FrameType{id: "TPOS", description: "Part of a set", constructor: NewTextFrame},
		"TPUB": FrameType{id: "TPUB", description: "Publisher", constructor: NewTextFrame},
		"TRCK": FrameType{id: "TRCK", description: "Track number/Position in set", constructor: NewTextFrame},
		"TRDA": FrameType{id: "TRDA", description: "Recording dates", constructor: NewTextFrame},
		"TRSN": FrameType{id: "TRSN", description: "Internet radio station name", constructor: NewTextFrame},
		"TRSO": FrameType{id: "TRSO", description: "Internet radio station owner", constructor: NewTextFrame},
		"TSIZ": FrameType{id: "TSIZ", description: "Size", constructor: NewTextFrame},
		"TSRC": FrameType{id: "TSRC", description: "ISRC (international standard recording code)", constructor: NewTextFrame},
		"TSSE": FrameType{id: "TSSE", description: "Software/Hardware and settings used for encoding", constructor: NewTextFrame},
		"TYER": FrameType{id: "TYER", description: "Year", constructor: NewTextFrame},
		"TXXX": FrameType{id: "TXXX", description: "User defined text information frame", constructor: NewDescTextFrame},
		"UFID": FrameType{id: "UFID", description: "Unique file identifier", constructor: NewDataFrame},
		"USER": FrameType{id: "USER", description: "Terms of use", constructor: NewDataFrame},
		"USLT": FrameType{id: "USLT", description: "Unsychronized lyric/text transcription", constructor: NewUnsynchTextFrame},
		"WCOM": FrameType{id: "WCOM", description: "Commercial information", constructor: NewDataFrame},
		"WCOP": FrameType{id: "WCOP", description: "Copyright/Legal information", constructor: NewDataFrame},
		"WOAF": FrameType{id: "WOAF", description: "Official audio file webpage", constructor: NewDataFrame},
		"WOAR": FrameType{id: "WOAR", description: "Official artist/performer webpage", constructor: NewDataFrame},
		"WOAS": FrameType{id: "WOAS", description: "Official audio source webpage", constructor: NewDataFrame},
		"WORS": FrameType{id: "WORS", description: "Official internet radio station homepage", constructor: NewDataFrame},
		"WPAY": FrameType{id: "WPAY", description: "Payment", constructor: NewDataFrame},
		"WPUB": FrameType{id: "WPUB", description: "Publishers official webpage", constructor: NewDataFrame},
		"WXXX": FrameType{id: "WXXX", description: "User defined URL link frame", constructor: NewDataFrame},
	}
)
