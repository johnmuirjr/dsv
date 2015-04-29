Title: dsv Package Readme
Author: Jordan Vaughan
Date: April 29, 2015
Format: Plain text
Encoding: UTF-8

INTRODUCTION
    The dsv package provides Reader and Writer types and associated functions
    for reading and writing records in delimiter-separated values (DSV)
    format.  The format is described in detail in chapter five, “Textuality”,
    of Eric Steven Raymond, “The Art of Unix Programming” (Boston: Addison-
    Wesley), 2003.  It’s much like comma-separated values (CSV) but is much
    easier to write and parse.

    Here are some example DSV records (using colon characters (':') and
    reverse solidus characters ('\') as field separator and escape
    characters, respectively) and their CSV equivalents:

        DSV
          "No man is free who cannot command himself.":Pythagoras:Greek
        CSV
          """No man is free who cannot command himself.""",Pythagoras,Greek

        DSV
          Consider this\: What is the best use of one's life?:Me
        CSV
          Consider this: What is the best use of one's life?,Me

        DSV
          Red pajamas, yellow pajamas, brown pajamas.:Japanese tongue twister
        CSV
          "Red pajamas, yellow pajamas, brown pajamas.",Japanese tongue twister

        DSV
          1:2.5:-5:0
          1,2.5,-5,0

    The interface is similar to the standard csv package’s.  See MANUAL.txt
    for details or run godoc after downloading the source code.

COPYRIGHT
    Copyright?  Hah!  Here's my “copyright”:

        This package was written in 2015 by Jordan Vaughan.

        To the extent possible under law, the author(s) have dedicated all
        copyright and related and neighboring rights to this software to the
        public domain worldwide. This software is distributed without any
        warranty.

        You should have received a copy of the CC0 Public Domain Dedication
        along with this software. If not, see
        http://creativecommons.org/publicdomain/zero/1.0/.

    A full copy of the public domain dedication is in COPYING.
