# Open Quiz Format for Typst

The Open Quiz Format (OQF) is a markup language for creating multiple choice
quizzes. It attempts to be as simple to use as possible for anyone to use.

## Format

A `.oqf` file can be made in any text editor and has an extremely simple
format. Questions and answers are formatted with the following syntax:

```oqf
:Question
!Incorrect answers are marked with a ! and not a =
=Correct answers are marked with a = and not a !
=Multiple correct answers can be selected at the same time
!Questions can have as many questions as you want
=Questions end with a semi colon character;
```

See the GitHub Repo for more detailed information

## Additional Notes

This is the reference implementation for OQF. If one is looking to create a
library for OQF, use this as a reference
