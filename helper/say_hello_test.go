package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//Before
	fmt.Println("Before Unit Test")

	m.Run()

	//After
	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("Nur", func(t *testing.T) {
		result := SayHello("dayat")
		require.Equal(t, "Hello dayat", result, "result must be 'Hello dayat'")
	})
	t.Run("Hidayat", func(t *testing.T) {
		result := SayHello("dayato")
		require.Equal(t, "Hello dayato", result, "result must be 'Hello dayato'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac os bro")
	}

	result := SayHello("dayat")
	require.Equal(t, "Hello dayat", result, "result must be 'Hello dayat'")
}

func TestSayHelloRequire(t *testing.T) {
	result := SayHello("dayat")
	require.Equal(t, "Hello dayat", result, "result must be 'Hello dayat'")
	fmt.Println("TestSayHelloAssert with Assert Done")
}
func TestSayHelloAssert(t *testing.T) {
	result := SayHello("dayat")
	assert.Equal(t, "Hello dayat", result, "result must be 'Hello dayat'")
	fmt.Println("TestSayHelloAssert with Assert Done")
}

func TestSayHello(t *testing.T) {
	result := SayHello("dayat")

	if result != "Hello dayat" {
		// error
		t.Error("Result is not 'Hello dayat'")
	}

	fmt.Println("halo bos")

}
func TestSayHelloccnt(t *testing.T) {
	result := SayHello("ccnt")

	if result != "Hello ccnt" {
		// error
		t.Fatal("Result is not 'Hello ccnt'")
	}

	fmt.Println("halo bosku")

}
