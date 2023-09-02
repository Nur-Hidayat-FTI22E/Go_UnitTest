package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "Nur",
			request: "Nur",
		},
		{
			name:    "Hidayat",
			request: "Hidayat",
		},
	}

	for _, bebenchmark := range benchmark {
		b.Run(bebenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(bebenchmark.request)
			}
		})
	}
}

func BenchmarkSubBenchmark(b *testing.B) {
	b.Run("Nur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Nur")
		}
	})
	b.Run("Hidayat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Hidayat")
		}
	})
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("dayat")
	}
}

func TestTableSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Nur",
			request:  "Nur",
			expected: "Hello Nur",
		},
		{
			name:     "Hidayat",
			request:  "Hidayat",
			expected: "Hello Hidayat",
		},
		{
			name:     "Coconut",
			request:  "Coconut",
			expected: "Hello Coconut",
		},
	}

	for _, tests := range tests {
		t.Run(tests.name, func(t *testing.T) {
			result := SayHello(tests.request)
			require.Equal(t, tests.expected, result)
		})
	}
}

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
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux bro")
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
