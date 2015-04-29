
package dsv

import (
    "bytes"
    "fmt"
    "strings"
    "testing"
)

func TestDSV(t *testing.T) {
    input := ":a: :a : a::\\::\\\n\nThis:is:a:\"test\\\\"
    expectedOutput := [][]string {
        {"", "a", " ", "a ", " a", "", ":", "\n"},
        {"This", "is", "a", "\"test\\"},
    }

    reader := NewReader(strings.NewReader(input))
    output, err := reader.ReadAll()
    if err != nil {
        t.Fatal("error while reading valid DSV string")
    }
    t.Logf(fmt.Sprintf("%v", output))
    if len(output) != len(expectedOutput) {
        t.Fatal(fmt.Sprintf("output doesn't have the expected number of records: %v instead of %v",
            len(output), len(expectedOutput)))
    }
    for n, result := range output {
        if len(result) != len(expectedOutput[n]) {
            t.Fatal(fmt.Sprintf("output record %v doesn't have same length as expected record: %v instead of %v",
                n, len(result), len(expectedOutput[n])))
        }
        for m, str := range result {
            if str != expectedOutput[n][m] {
                t.Fatal("output field isn't expected field")
            }
        }
    }

    buffer := bytes.Buffer{}
    writer := NewWriter(&buffer)
    err = writer.WriteAll(output)
    if err != nil {
        t.Fatal("error while writing DSV fields")
    }
    encoded := buffer.String()
    if encoded != input + "\n" {
        t.Fatal("written DSV doesn't match original DSV string")
    }
}
