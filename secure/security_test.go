package secure

import (
	"testing"
	"time"
)

var sha256Map = map[string]string{
	"helloWorld":   "11d4ddc357e0822968dbfd226b6e1c2aac018d076a54da4f65e1dc8180684ac3",
	"pengtaiKorea": "0ac49c86ec2dc8ef7a5b5ef3b0d7456bdf745cd25b721d0218e9c0153587f474",
	"한국어":          "ea3252281bc3bcec5672d33ec8c25dee8abd77235f85b8a049948f59ddf60b8d",
	"펑타이코리아":       "5ae860a64108311a40b3e4894e9435d83794dfb4751f2aa332069dd63c962b60",
}

func TestSha256(t *testing.T) {
	for o, h := range sha256Map {
		if h != Sha256(o) {
			t.Errorf("sha256 unmatched %s: expected %s but %s", o, h, Sha256(o))
		}
	}
}

func TestSha512(t *testing.T) {
	for o := range sha256Map {
		t.Logf("%s => %s", o, Sha512(o))
	}
}

func TestGenerateToken(t *testing.T) {
	testSize := 100000
	stack := make([]string, testSize)

	for i := 0; i < testSize; i++ {
		stack[i] = GenerateRandom512()
	}

	for i := 0; i < testSize; i++ {
		for j := 0; j < testSize; j++ {
			if i == j {
				continue
			}
			if stack[i] == stack[j] {
				t.Errorf("unexpected match %d %d", i, j)
			}
		}
	}
}

func TestKeygen(t *testing.T) {
	testSize := 100000
	stack := make([]string, testSize)

	for i := 0; i < testSize; i++ {
		stack[i] = GenerateKey()
		// fmt.Println("- " + stack[i])
		time.Sleep(10 * time.Millisecond)
	}

	for i := 0; i < testSize; i++ {
		for j := 0; j < testSize; j++ {
			if i == j {
				continue
			}
			if stack[i] == stack[j] {
				t.Errorf("generated key matched: %d vs. %d %s", i, j, stack[i])
			}
		}
	}
}

func TestEncryption(t *testing.T) {
	key := GenerateKey()

	origins := []string{
		"helloworld",
		"hello World",
		"PengtaiKorea",
		"PTK",
		"한국어",
		"펑타이",
		"펑타이코리아",
	}

	ciphers := make([]string, len(origins))

	for i, org := range origins {
		ciphers[i] = AESEncrypt(key, org)
		t.Logf(" - %s >> %s", org, ciphers[i])
	}

	// Right key decrypt
	for i, cy := range ciphers {
		if dec := AESDecrypt(key, cy); origins[i] != dec {
			t.Errorf("decrypt not matched %d: expect %s but %s", i, origins[i], dec)
		}
	}

	wkey := GenerateKey()
	// Wrong key decrypt
	for i, cy := range ciphers {
		if dec := AESDecrypt(wkey, cy); origins[i] == dec {
			t.Errorf("Expected wrong but hacked: %s (basekey %s, but %s)", origins[i], key, wkey)
		}
	}
}
