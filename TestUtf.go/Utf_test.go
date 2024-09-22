// Функция GetUTFLength(input []byte) (int, error)
// возвращает длину строки UTF8 и ошибку ErrInvalidUTF8 (в случае возникновения).
// Напишите тест, который бы проверял возвращаемые функцией значения.

// var ErrInvalidUTF8 = errors.New("invalid utf8")

// func GetUTFLength(input []byte) (int, error) {
// if !utf8.Valid(input) {
// return 0, ErrInvalidUTF8
// }

//  return utf8.RuneCount(input), nil
// }

package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name       string
		input      []byte
		wantLength int
		wantError  error
	}{
		{
			name:       "Valid UTF-8 string",
			input:      []byte("hello"),
			wantLength: 5,
			wantError:  nil,
		},
		{
			name:       "Invalid UTF-8 string",
			input:      []byte{0xff, 0xfe, 0xfd},
			wantLength: 0,
			wantError:  ErrInvalidUTF8,
		},
	}

	for _, tc := range cases {
		tc := tc // захват текущего значения tc
		t.Run(tc.name, func(t *testing.T) {
			gotLength, gotError := GetUTFLength(tc.input)

			// Проверяем длину
			if gotLength != tc.wantLength {
				t.Errorf("GetUTFLength(%v) = %v, want %v", tc.input, gotLength, tc.wantLength)
			}

			// Проверяем ошибку
			if gotError != tc.wantError {
				t.Errorf("GetUTFLength(%v) error = %v, want %v", tc.input, gotError, tc.wantError)
			}
		})
	}
}
