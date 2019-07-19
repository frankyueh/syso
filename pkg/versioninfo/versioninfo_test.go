package versioninfo

import (
	"testing"
)

func TestFormatVersionString(t *testing.T) {
	if formatVersionString(0x0001000200030004) != "1.2.3.4" {
		t.Errorf("failed")
	}
	vi := &VersionInfo{}
	if err := vi.SetFileVersionString("1.2.3.4"); err != nil {
		t.Fatal(err)
	}
	if vi.FileVersionString() != "1.2.3.4" {
		t.Errorf("failed")
	}
}

func TestParseVersionString(t *testing.T) {
	v, err := parseVersionString("1.2.3.4")
	if err != nil {
		t.Fatal(err)
	}
	if v != 0x0001000200030004 {
		t.Errorf("mismatching version; expected 0x0001000200030004, got %#016x", v)
	}
}

func TestString(t *testing.T) {
	vi := &VersionInfo{}
	vi.SetString(0x0409, 0x04b0, "foo", "bar")
	if s, ok := vi.String(0x0409, 0x04b0, "foo"); !ok {
		t.Fatal("cannot get string")
	} else if s != "bar" {
		t.Fatal("wrong string")
	}
	if _, ok := vi.String(0x1, 0x2, "foo"); ok {
		t.Fatal("must not get string")
	}
}

func TestFreezeEmpty(t *testing.T) {
	vi := &VersionInfo{}
	vi.freeze()
	if vi.length != 88 {
		t.Fatalf("wrong VersionInfo.length; expected 88, got %d", vi.length)
	}
	if vi.valueLength != 52 {
		t.Fatalf("wrong VersionInfo.valueLength; expected 52, got %d", vi.valueLength)
	}
}

func TestFreeze(t *testing.T) {
	vi := &VersionInfo{}
	vi.SetString(0x0409, 0x04b0, "foo", "bar")
	vi.freeze()
	if vi.length != 166 {
		t.Fatalf("wrong VersionInfo.length; expected 166, got %d", vi.length)
	}
	if vi.stringFileInfo.length != 78 {
		t.Fatalf("wrong VersionInfo.stringFileInfo.length; expected 78, got %d", vi.stringFileInfo.length)
	}
	if vi.stringFileInfo.stringTables[0].length != 42 {
		t.Fatalf("wrong VersionInfo.stringFileInfo.stringTables[0].length; expected 42, got %d", vi.stringFileInfo.stringTables[0].length)
	}
	if vi.stringFileInfo.stringTables[0].strings[0].length != 18 {
		t.Fatalf("wrong VersionInfo.stringFileInfo.stringTables[0].strings[0].length; expected 18, got %d", vi.stringFileInfo.stringTables[0].strings[0].length)
	}
	if vi.stringFileInfo.stringTables[0].strings[0].valueLength != 6 {
		t.Fatalf("wrong VersionInfo.stringFileInfo.stringTables[0].strings[0].valueLength; expected 6, got %d", vi.stringFileInfo.stringTables[0].strings[0].valueLength)
	}
}
