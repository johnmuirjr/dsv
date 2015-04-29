
// Package dsv reads and writes delimiter-separated values (DSV) files.
// Their format is described in detail in chapter five, “Textuality”,
// of Eric Steven Raymond, “The Art of Unix Programming” (Boston: Addison-
// Wesley), 2003.
//
// A DSV file contains zero or more records consisting of one or more fields
// per record.  Each field within a record is separated by a single character,
// which is usually the colon (':').  Records are separated from each other
// by one or more consecutive newline characters ('\n').  Any character
// (including separator characters and newlines) may be escaped by prefixing
// it with a single reverse solidus ('\\').  Whitespace is preserved within
// fields.  The final record may be optionally followed by one or more
// newline characters.
package dsv

import (
    "bufio"
    "bytes"
    "io"
)

// A Reader reads records from a DSV file.
//
// Readers returned by NewReader use reverse solidus characters ('\\') and
// colon characters (':') as escape and record separator characters,
// respectively.  The Reader's exported fields can be modified to change
// these settings.
type Reader struct {
    Escape      rune    // prefix for escaping characters
    Separator   rune    // field delimiter/separator
    reader      io.RuneReader
    field       bytes.Buffer
}

// A Writer writes records to an io.Writer in DSV format.
//
// Writers returned by NewWriter use reverse solidus characters ('\\') and
// colon characters (':') as escape and record separator characters,
// respectively.  The Writer's exported fields can be modified to change
// these settings.
type Writer struct {
    Escape      rune    // prefix for escaping characters
    Separator   rune    // field delimiter/separator
    writer      *bufio.Writer
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.RuneReader) *Reader {
    return &Reader {
        Escape:    '\\',
        Separator: ':',
        reader:    r,
    }
}

// Read reads one record from r.  The record is a slice of strings with each
// string representing one field.  err is nil if no errors occur or EOF is
// reached.  (EOF is not treated as an error.)
func (r *Reader) Read() (fields []string, err error) {
    var c rune
    var isEscaping bool

    // Eliminate leading newlines.
    for {
        c, _, err = r.reader.ReadRune()
        if err == io.EOF {
            return nil, nil
        }
        if err != nil {
            return nil, err
        }
        if c != '\n' {
            break
        }
    }

    defer r.field.Reset()

    // Parse the record (all fields up to the first unescaped newline).
    for {
        if isEscaping {
            r.field.WriteRune(c)
            isEscaping = false
        } else {
            switch c {
                case r.Separator:
                    fields = append(fields, r.field.String())
                    r.field.Reset()
                case r.Escape:
                    isEscaping = true
                case '\n':
                    fields = append(fields, r.field.String())
                    return fields, nil
                default:
                    r.field.WriteRune(c)
            }
        }
        c, _, err = r.reader.ReadRune()
        if err == io.EOF {
            fields = append(fields, r.field.String())
            break
        }
        if err != nil {
            fields = append(fields, r.field.String())
            break
        }
    }
    return
}

// ReadAll reads all remaining records from r.  Each record is a slice of
// fields, one string per field.  err is set to nil if no errors occur or
// EOF is reached.  (EOF is not treated as an error.)
func (r *Reader) ReadAll() (records [][]string, err error) {
    for {
        record, err := r.Read()
        if err == io.EOF {
            return append(records, record), nil
        }
        if err != nil {
            return nil, err
        }
        if record == nil {
            return records, nil
        }
        records = append(records, record)
    }
}

// NewWriter returns a Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
    return &Writer {
        Escape:    '\\',
        Separator: ':',
        writer:    bufio.NewWriter(w),
    }
}

// Error reports any error that occurred during the last Flush or Write.
func (w *Writer) Error() error {
    _, err := w.writer.Write(nil)
    return err
}

// Flush writes buffered data to w's underlying io.Writer.  Call Error to
// check for errors.
func (w *Writer) Flush() {
    w.writer.Flush()
}

// Write writes a single record to w.  The record is a slice of strings
// representing its fields, one string per field.  Characters within the
// fields are escaped as necessary.
func (w *Writer) Write(record []string) (err error) {
    for n, field := range record {
        if n > 0 {
            if _, err = w.writer.WriteRune(w.Separator); err != nil {
                return
            }
        }
        for _, r := range field {
            switch r {
                case w.Escape:
                    _, err = w.writer.WriteRune(w.Escape)
                    if err == nil {
                        _, err = w.writer.WriteRune(w.Escape)
                    }
                case w.Separator:
                    _, err = w.writer.WriteRune(w.Escape)
                    if err == nil {
                        _, err = w.writer.WriteRune(w.Separator)
                    }
                case '\n':
                    _, err = w.writer.WriteRune(w.Escape)
                    if err == nil {
                        err = w.writer.WriteByte('\n')
                    }
                default:
                    _, err = w.writer.WriteRune(r)
            }
            if err != nil {
                return
            }
        }
    }
    err = w.writer.WriteByte('\n')
    return
}

// WriteAll writes multiple records to w and calls Flush.
func (w *Writer) WriteAll(records [][]string) (err error) {
    for _, record := range records {
        if err = w.Write(record); err != nil {
            return
        }
    }
    return w.writer.Flush()
}
