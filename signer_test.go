package jwt

import "testing"

var key = []byte("private")

func TestHMACSigner(t *testing.T) {
	var tests = []struct {
		in  []byte
		out []byte
	}{
		{
			[]byte("foo"),
			[]byte{
				0x87, 0x5c, 0xb3, 0xc1, 0x8b, 0xa6, 0xb7, 0x55, 0x97, 0x24,
				0xe6, 0x07, 0x3b, 0xd0, 0x81, 0x64, 0xe9, 0x0d, 0xea, 0x07,
				0x6d, 0xa0, 0x24, 0x32, 0xed, 0x4e, 0xd6, 0x2a, 0x9c, 0x44,
				0x94, 0xdb,
			},
		},
		{
			[]byte("bar"),
			[]byte{
				0x65, 0xc8, 0x5b, 0x0d, 0xfd, 0x14, 0xf8, 0x65, 0x95, 0x3f,
				0xde, 0x63, 0x38, 0xcb, 0xe7, 0xbd, 0xdc, 0x56, 0x29, 0x86,
				0xe0, 0xe6, 0x43, 0xe0, 0x5d, 0x93, 0x18, 0xff, 0x2c, 0xa2,
				0xce, 0x99,
			},
		},
		{
			[]byte("baz"),
			[]byte{
				0xcd, 0xbb, 0xdd, 0x4c, 0xe2, 0xf6, 0xbd, 0xfb, 0xf0, 0x10,
				0x2a, 0xe0, 0x5a, 0x0c, 0xf4, 0xa2, 0xb9, 0x7a, 0x57, 0x48,
				0x38, 0x02, 0x33, 0x23, 0x87, 0x9d, 0x74, 0x73, 0xa4, 0x05,
				0x9d, 0x9c,
			},
		},
	}
	for i, tt := range tests {
		out, err := HS256.Sign(tt.in, key)
		if err != nil {
			t.Errorf("%d. Sign err\nhave %v\nwant %v", i, err, nil)
			continue
		}
		err = HS256.Verify(tt.in, tt.out, key)
		if err != nil {
			t.Errorf("%d. Verify\nhave %v\n     % #010x\nwant %v", i, err, out, nil)
		}
	}
}
