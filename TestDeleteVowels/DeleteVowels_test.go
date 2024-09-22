// Напишите тест для функции DeleteVowels(s string) string,
// которая должна удалять все гласные из строки английского языка (y не считается гласной).
// Используйте table driven testing.

package main

import (
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name   string
		values string
		want   string
	}{
		{
			name:   "with vowels",
			values: "abcdefghijklmnopqrstuvwxyz",
			want:   "bcdfghjklmnpqrstvwxyz",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteVowels(tc.values)

			if got != tc.want {
				t.Errorf("DeleteVowels(%v) = %v, want %v", tc.values, got, tc.want)
			}
		})
	}
}
