package mymode2

import "testing"

func TestHelloMode2(t *testing.T) {
	want := "Hello Mode2"
	if got := HelloMode2(); want != got {
		t.Errorf("HelloMode2() = %q, want %q", got, want)
	}
}

func TestPlayMusic(t *testing.T) {
	want := "Mode2 playing music..."
	if got := PlayMusic(); want != got {
		t.Errorf("TestPlayMusic() = %q, want %q", got, want)
	}
}



func TestSayHello(t *testing.T) {
	name := "zs"
	want := "Hello," + name
	if got:= SayHello(name); want!= got {
		t.Errorf("SayHello(" + name + ") = %q, want %q", got, want)
	}
}

func TestWatchmovie(t *testing.T) {
	want := "Mode2 watchmovie..."
	if got := watchmovie(); want != got {
		t.Errorf("watchmovie() = %q, want %q", got, want)
	}
}
