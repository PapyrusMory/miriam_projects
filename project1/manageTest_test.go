package main

import (
	"reflect"
	"testing"
)

func TestRemoveIndex(t *testing.T){
	a := []string{"Hello", ",","World!"}
	b:= []string{"Hello","World!"}
	v := RemoveIndex(a,1)

	if reflect.DeepEqual(v, b) != true {
		t.Error("Expected true, got ", v)
	}
}

func TestHexToDeci(t *testing.T){
	
	b := []string{"30", "files","seerr"}
	a := []string{"1E","(hex)", "files","seerr"}
	v := HexToDeci(a,1)
	if reflect.DeepEqual(RemoveDuplicate(v),b) != true{
		t.Error("Expected true, got ", RemoveDuplicate(v))
	}
	
}

func TestRemoveDuplicate(t *testing.T){
	a:= []string{"Bonjour,", "j'espère", "que", "vous", "allez", "bien.", "bien.", "bien."}
	b:= []string{"Bonjour,", "j'espère", "que", "vous", "allez", "bien."}
	v := RemoveDuplicate(a)

	if reflect.DeepEqual(v,b) !=true {
		t.Error("Expected true, got", v)
	}
}

func TestManagePunc(t *testing.T){
	a:= []string{"Bonjour","Monsieurs",",","soyez", "les", "bienvenus."}
	b := []string{"Bonjour","Monsieurs,","soyez", "les", "bienvenus."}
	v := ManagePunc(a,2)

	if reflect.DeepEqual(RemoveDuplicate(v), b) != true{
		t.Error("Expecteed true, got ", RemoveDuplicate(v))
	}
}

func TestGetUpNumFunc(t *testing.T){
	a := []string{"Il", "est", "minuit", "et", "(up,3)", "je", "code"}
	b := []string{"Il", "EST", "MINUIT", "ET", "je", "code"}

	v := GetUpNumFunc("(up,3)", a, 4)

	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestGetCapNumFunc(t *testing.T){
	a := []string{"Il", "est", "minuit", "et", "(cap,3)", "je", "code"}
	b := []string{"Il", "Est", "Minuit", "Et", "je", "code"}

	v := GetCapNumFunc("(cap,3)", a, 4)

	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestGetLowNumFunc(t *testing.T){
	a := []string{"Il", "EST", "MINUIT", "ET", "(low,3)", "je", "code"}
	b := []string{"Il", "est", "minuit", "et", "je", "code"}

	v := GetLowNumFunc("(low,3)", a, 4)

	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestBinFunc(t *testing.T){
	a := []string{"It", "has", "been", "10", "(bin)", "years"}
	b := []string{"It", "has", "been", "2", "years"}
	v := BinFunc(a,4)
	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestUpFunc(t *testing.T){
	a := []string{"I", "am", "happy", "(up)", "to", "see", "you"}
	b := []string{"I", "am", "HAPPY", "to", "see", "you"}
	v := UpFunc(a,3)
	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestCapFunc(t *testing.T){
	a := []string{"I", "am", "happy", "(up)", "to", "see", "you"}
	b := []string{"I", "am", "Happy", "to", "see", "you"}
	v := CapFunc(a,3)
	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestLowFunc(t *testing.T){
	a := []string{"I", "am", "HAPPY", "(low)", "to", "see", "you"}
	b := []string{"I", "am", "happy", "to", "see", "you"}
	v := LowFunc(a,3)
	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestAddNFunc(t *testing.T){
	a := []string{"A", "amazing", "thing", "it's", "over"}
	b := []string{"An", "amazing", "thing", "it's", "over"}
	v := AddNFunc(a,0)
	if reflect.DeepEqual(RemoveDuplicate(v), b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
	}
}

func TestPuncFunc(t *testing.T){
	a := []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)" , ".Yes", "I", "(cap)", "..."}
	b := []string{"Welcome", "to", "the", "Brooklyn", "bridge.", "(cap)", "Yes","I...", "(cap)", "..."}
	
	v := PuncFunc(a,6)
	r := PuncFunc(a, 9)
	if reflect.DeepEqual(v, b) != true && reflect.DeepEqual(r,b) != true {
		t.Error("Expected true, got", v)
		t.Error("Expected true, got", r)
	}
}

func TestManageAp(t *testing.T){
	a := []string{"A","'", "amazing","'", "thing","'", "its", "over","'", "Good"}
	b := []string{"A", "'amazing'", "thing", "'its", "over'", "Good"}
	r := ManageAp(a, 5)
	v := ManageAp(a,1)
	
	if reflect.DeepEqual(RemoveDuplicate(v), b) != true && reflect.DeepEqual(RemoveDuplicate(r),b) != true {
		t.Error("Expected true, got", RemoveDuplicate(v))
		t.Error("Expected true, got", RemoveDuplicate(r))
	}
}