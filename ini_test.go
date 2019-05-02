package ini

import (
	"fmt"
	"testing"
)

const (
	inFileName  = "testin.ini"
	outFileName = "testout.ini"
)

func Test_tKeyVal_String(t *testing.T) {
	kv1 := TKeyVal{"key1", "val1"}
	rs1 := "key1 = val1"
	kv2 := TKeyVal{"key2", ""}
	rs2 := "key2 ="
	kv3 := TKeyVal{"", ""}
	rs3 := " ="
	type fields struct {
		Key   string
		Value string
	}
	tests := []struct {
		name string
		kv   *TKeyVal
		want string
	}{
		// TODO: Add test cases.
		{" 1", &kv1, rs1},
		{" 2", &kv2, rs2},
		{" 3", &kv3, rs3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := TKeyVal{
				Key:   tt.kv.Key,
				Value: tt.kv.Value,
			}
			if got := kv.String(); got != tt.want {
				t.Errorf("tKeyVal.String() = {%v}, want {%v}", got, tt.want)
			}
		})
	}
} // Test_tKeyVal_String()

func Benchmark_tKeyVal_String(b *testing.B) {
	kv1 := TKeyVal{"key1", "val1"}
	for n := 0; n < b.N; n++ {
		if 0 > len(kv1.String()) {
			continue
		}
	}
} // Benchmark_tKeyVal_String()

func Benchmark_tKeyVal_string0(b *testing.B) {
	kv1 := TKeyVal{"key1", "val1"}
	for n := 0; n < b.N; n++ {
		if 0 > len(kv1.string0()) {
			continue
		}
	}
} // Benchmark_tKeyVal_string0()

func Test_trimRemoveQuotes(t *testing.T) {
	si1 := "'this is a text'"
	so1 := "this is a text"
	si2 := " \" this is a text \" "
	so2 := " this is a text "
	si3 := " \" this is a text ' "
	so3 := "\" this is a text '"
	type args struct {
		aString string
	}
	tests := []struct {
		name          string
		args          args
		wantRFiltered string
	}{
		// TODO: Add test cases.
		{" 1", args{si1}, so1},
		{" 2", args{si2}, so2},
		{" 3", args{si3}, so3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRFiltered := trimRemoveQuotes(tt.args.aString); gotRFiltered != tt.wantRFiltered {
				t.Errorf("trimRemoveQuotes() =\n{%v}, want {%v}", gotRFiltered, tt.wantRFiltered)
			}
		})
	}
} // Test_trimRemoveQuotes()

func Test_tSection_String(t *testing.T) {
	sl1 := TSection{
		TKeyVal{"key1", "val1"},
		TKeyVal{"key2", "val2"},
		TKeyVal{"key3", "val3"},
		TKeyVal{"key4", ""},
	}
	rl1 := "key1 = val1\nkey2 = val2\nkey3 = val3\nkey4 =\n"
	tests := []struct {
		name string
		cs   *TSection
		want string
	}{
		// TODO: Add test cases.
		{" 1", &sl1, rl1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.String(); got != tt.want {
				t.Errorf("tSection.String() = {%v}, want {%v}", got, tt.want)
			}
		})
	}
} // Test_tSection_String()

func Benchmark_TSection_String(b *testing.B) {
	sl1 := TSection{
		TKeyVal{"key1", "val1"},
		TKeyVal{"key2", "val2"},
		TKeyVal{"key3", "val3"},
		TKeyVal{"key4", ""},
	}
	for n := 0; n < b.N; n++ {
		if 0 > len(sl1.String()) {
			continue
		}
	}
} // Benchmark_TSection_String()

func Benchmark_TSection_string0(b *testing.B) {
	sl1 := TSection{
		TKeyVal{"key1", "val1"},
		TKeyVal{"key2", "val2"},
		TKeyVal{"key3", "val3"},
		TKeyVal{"key4", ""},
	}
	for n := 0; n < b.N; n++ {
		if 0 > len(sl1.string0()) {
			continue
		}
	}
} // Benchmark_TSection_string0()

func TestIniSections_Clear(t *testing.T) {
	cis, _ := LoadFile(inFileName)
	type fields struct {
		defSect  string
		secOrder []string
		sections tIniSections
	}
	cs := fields(*cis)
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"1", cs, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &TSections{
				defSect:  tt.fields.defSect,
				secOrder: tt.fields.secOrder,
				sections: tt.fields.sections,
			}
			if got := id.Clear(); got != tt.want {
				t.Errorf("IniSections.Clear() = '%v', want '%v'", got, tt.want)
			}
		})
	}
} // TestIniSections_Clear()

func TestIniSections_RemoveSection(t *testing.T) {
	cis, _ := LoadFile(inFileName)
	type fields struct {
		defSect  string
		secOrder []string
		sections tIniSections
	}
	type args struct {
		aSection string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"1", fields(*cis), args{"Default"}, true}, // first
		{"2", fields(*cis), args{"port3"}, true},   // last
		{"3", fields(*cis), args{"sql3"}, true},    // middle
		{"4", fields(*cis), args{"nichda"}, true},  // n.a.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &TSections{
				defSect:  tt.fields.defSect,
				secOrder: tt.fields.secOrder,
				sections: tt.fields.sections,
			}
			if got := id.RemoveSection(tt.args.aSection); got != tt.want {
				t.Errorf("IniSections.RemoveSection() = '%v', want '%v'", got, tt.want)
			}
		})
	}
} // TestIniSections_RemoveSection()

func TestIniSections_RemoveSectionKey(t *testing.T) {
	cis, _ := LoadFile(inFileName)
	type fields struct {
		defSect  string
		secOrder tOrder
		sections tIniSections
	}
	type args struct {
		aSection string
		aKey     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"1", fields(*cis), args{"", "ach jeh"}, true},
		{"2", fields(*cis), args{"Default", "ach jeh"}, true},
		{"3", fields(*cis), args{"sql3", "password"}, true},
		{"4", fields(*cis), args{"sql3", "port"}, true},
		{"5", fields(*cis), args{"port0", ""}, true},
		{"6", fields(*cis), args{"", ""}, true},
		{"7", fields(*cis), args{"general", "n.a."}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &TSections{
				defSect:  tt.fields.defSect,
				secOrder: tt.fields.secOrder,
				sections: tt.fields.sections,
			}
			if got := id.RemoveSectionKey(tt.args.aSection, tt.args.aKey); got != tt.want {
				t.Errorf("IniSections.RemoveSectionKey() = '%v', want '%v'", got, tt.want)
			}
		})
	}
} // TestIniSections_RemoveSectionKey()

func TestIniSections_String(t *testing.T) {
	id1 := TSections{
		defSect: "Default",
		secOrder: tOrder{
			"Default",
			"Sect2",
			"NOOP",
		},
		sections: tIniSections{
			"Sect2": &TSection{
				TKeyVal{"key3", "val3"},
				TKeyVal{"key4", ""},
			},
			"Default": &TSection{
				TKeyVal{"key1", "val1"},
				TKeyVal{"key2", "val2"},
			},
		},
	}
	rl1 := "\n[Default]\nkey1 = val1\nkey2 = val2\n\n[Sect2]\nkey3 = val3\nkey4 =\n"
	type fields struct {
		defSect  string
		secOrder tOrder
		Sections tIniSections
	}
	tests := []struct {
		name   string
		fields TSections
		want   string
	}{
		// TODO: Add test cases.
		{" 1", id1, rl1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &TSections{
				defSect:  tt.fields.defSect,
				secOrder: tt.fields.secOrder,
				sections: tt.fields.sections,
			}
			if got := id.String(); got != tt.want {
				t.Errorf("IniSections.String() =\n{%v}, want\n{%v}", got, tt.want)
			}
		})
	}
} // TestIniSections_String()

func Benchmark_TSections_String(b *testing.B) {
	id1 := TSections{
		defSect: "Default",
		secOrder: tOrder{
			"Default",
			"Sect2",
			"NOOP",
		},
		sections: tIniSections{
			"Sect2": &TSection{
				TKeyVal{"key3", "val3"},
				TKeyVal{"key4", ""},
			},
			"Default": &TSection{
				TKeyVal{"key1", "val1"},
				TKeyVal{"key2", "val2"},
			},
		},
	}
	for n := 0; n < b.N; n++ {
		if 0 > len(id1.String()) {
			continue
		}
	}
} // Benchmark_TSections_String()

func Benchmark_TSections_string0(b *testing.B) {
	id1 := TSections{
		defSect: "Default",
		secOrder: tOrder{
			"Default",
			"Sect2",
			"NOOP",
		},
		sections: tIniSections{
			"Sect2": &TSection{
				TKeyVal{"key3", "val3"},
				TKeyVal{"key4", ""},
			},
			"Default": &TSection{
				TKeyVal{"key1", "val1"},
				TKeyVal{"key2", "val2"},
			},
		},
	}
	for n := 0; n < b.N; n++ {
		if 0 > len(id1.string0()) {
			continue
		}
	}
} // Benchmark_TSections_string0()

func compare1(aString string) {
	if "" == aString {
		return
	}
} // compare1()

func compare2(aString string) {
	if 0 == len(aString) {
		return
	}
} // compare2()

func Benchmark_compare1(b *testing.B) {
	for n := 0; n < b.N*8; n++ {
		compare1("qwertzuiopü+#äölkjhgfdsa<yxcvbnm,.-^1234567890ß´qwertzuiop")
	}
} // Benchmark_compare1()

func Benchmark_compare2(b *testing.B) {
	for n := 0; n < b.N*8; n++ {
		compare2("qwertzuiopü+#äölkjhgfdsa<yxcvbnm,.-^1234567890ß´qwertzuiop")
	}
} // Benchmark_compare2()

func TestIniSections_updateSectKey(t *testing.T) {
	cis, _ := LoadFile(inFileName)
	type fields struct {
		defSect  string
		secOrder []string
		sections tIniSections
	}
	cs := fields(*cis)
	type args struct {
		aSection string
		aKey     string
		aValue   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"1", cs, args{"", "", ""}, false},
		{"2", cs, args{"general", "", ""}, false},
		{"3", cs, args{"", "loglevel", ""}, true},
		{"4", cs, args{"general", "loglevel", ""}, true},
		{"5", cs, args{"general", "loglevel", "8"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &TSections{
				defSect:  tt.fields.defSect,
				secOrder: tt.fields.secOrder,
				sections: tt.fields.sections,
			}
			if got := id.updateSectKey(tt.args.aSection, tt.args.aKey, tt.args.aValue); got != tt.want {
				t.Errorf("IniSections.updateSectKey() = '%v', want '%v'", got, tt.want)
			}
		})
	}
} // TestIniSections_updateSectKey()

func TestIniSections_UpdateSectKeyBool(t *testing.T) {
	cis, _ := LoadFile(inFileName)
	type fields struct {
		defSect  string
		secOrder []string
		sections tIniSections
	}
	cs := fields(*cis)
	type args struct {
		aSection string
		aKey     string
		aValue   bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"1", cs, args{"", "", false}, false},
		{"2", cs, args{"general", "", true}, false},
		{"3", cs, args{"", "loglevel", false}, true},
		{"4", cs, args{"general", "loglevel", true}, true},
		{"5", cs, args{"general", "loglevel", false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &TSections{
				defSect:  tt.fields.defSect,
				secOrder: tt.fields.secOrder,
				sections: tt.fields.sections,
			}
			if got := id.UpdateSectKeyBool(tt.args.aSection, tt.args.aKey, tt.args.aValue); got != tt.want {
				t.Errorf("IniSections.UpdateSectKeyBool() = '%v', want '%v'", got, tt.want)
			}
		})
	}
} // TestIniSections_UpdateSectKeyBool()

func TestIniSections_WriteFile(t *testing.T) {
	type args struct {
		aFilename string
	}
	id, _ := LoadFile(inFileName)
	tests := []struct {
		name       string
		id         *TSections
		args       args
		wantRBytes int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{"1",
			id,
			args{aFilename: outFileName},
			4253,
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRBytes, err := tt.id.Store(tt.args.aFilename)
			if (err != nil) != tt.wantErr {
				t.Errorf("IniSections.WriteFile() error = {%v}, wantErr {%v}", err, tt.wantErr)
				return
			}
			if gotRBytes != tt.wantRBytes {
				t.Errorf("IniSections.WriteFile() = '%v', want '%v'", gotRBytes, tt.wantRBytes)
			}
		})
	}
} // TestIniSections_WriteFile()

func walkFunc(aSect, aKey, aVal string) {
	fmt.Printf("\nSection: %s\nKey: %s\nValue: %s\n", aSect, aKey, aVal)
} // walkFunc()

func TestTSections_Walk(t *testing.T) {
	il, _ := LoadFile(inFileName)
	type args struct {
		aFunc TWalkFunc
	}
	tests := []struct {
		name   string
		fields TSections
		args   args
	}{
		// TODO: Add test cases.
		{" 1", *il, args{walkFunc}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Walk(tt.args.aFunc)
		})
	}
} // TestTSections_Walk()

type tTestWalk int

func (tw tTestWalk) Walk(aSect, aKey, aVal string) {
	fmt.Printf("\nSection: %s\nKey: %s\nValue: %s\n", aSect, aKey, aVal)
} // walkFunc()

func TestTSections_Walker(t *testing.T) {
	type args struct {
		aWalker TIniWalker
	}
	il, _ := LoadFile(inFileName)
	var walker tTestWalk
	tests := []struct {
		name   string
		fields *TSections
		args   args
	}{
		// TODO: Add test cases.
		{" 1", il, args{walker}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Walker(tt.args.aWalker)
		})
	}
} // TestTSections_Walker()

func TestTSections_RemoveSectionKey(t *testing.T) {
	type fields struct {
		defSect  string
		secOrder tOrder
		sections tIniSections
	}
	type args struct {
		aSection string
		aKey     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &TSections{
				defSect:  tt.fields.defSect,
				secOrder: tt.fields.secOrder,
				sections: tt.fields.sections,
			}
			if got := id.RemoveSectionKey(tt.args.aSection, tt.args.aKey); got != tt.want {
				t.Errorf("TSections.RemoveSectionKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTKeyVal_String(t *testing.T) {
	type fields struct {
		Key   string
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := &TKeyVal{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := kv.String(); got != tt.want {
				t.Errorf("TKeyVal.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTSection_String(t *testing.T) {
	tests := []struct {
		name        string
		cs          *TSection
		wantRString string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRString := tt.cs.String(); gotRString != tt.wantRString {
				t.Errorf("TSection.String() = %v, want %v", gotRString, tt.wantRString)
			}
		})
	}
}