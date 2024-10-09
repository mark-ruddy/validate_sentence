# Validate Sentence
This Go program can be used to validate a sentence.

## Usage
Compile binary and pass sentence as flag:

```
go build .
./validate_sentence -sentence "A correct sentence."
```

Enter sentence with user input and enable debug logs to view reason for sentence being invalid:

```
/validate_sentence --debug
Enter a sentence: "The quick brown fox said "hello Mr lazy dog."
DEBU[0017] Sentence has unequal quotes: '"The quick brown fox said "hello Mr lazy dog."'
Invalid sentence: '"The quick brown fox said "hello Mr lazy dog."'
```

## Testing
Run unit tests:

```
go test .
```

Run unit tests with verbose log output:

```
got test . -v
```

Example verbose test output:

```
go test . -v
=== RUN   TestValidSentences
--- PASS: TestValidSentences (0.00s)
=== RUN   TestInvalidSentences
time="2024-10-09T15:39:20+01:00" level=debug msg="Sentence has more than one period: 'The quick brown fox said \"hello Mr. lazy dog\".'"
time="2024-10-09T15:39:20+01:00" level=debug msg="Sentence does not start with an uppercase letter: 'the quick brown fox said \"hello Mr lazy dog\".'"
time="2024-10-09T15:39:20+01:00" level=debug msg="Sentence has unequal quotes: '\"The quick brown fox said \"hello Mr lazy dog.\"'"
time="2024-10-09T15:39:20+01:00" level=debug msg="Sentence does not contain spelled out numbers: 'One lazy dog is too few, 12 is too many.'"
time="2024-10-09T15:39:20+01:00" level=debug msg="Sentence does not contain spelled out numbers: 'Are there 11, 12, or 13 lazy dogs?'"
--- PASS: TestInvalidSentences (0.00s)
=== RUN   TestValidStartsWithUppercase
--- PASS: TestValidStartsWithUppercase (0.00s)
=== RUN   TestInvalidStartsWithUppercase
--- PASS: TestInvalidStartsWithUppercase (0.00s)
=== RUN   TestValidHasEqualQuotes
--- PASS: TestValidHasEqualQuotes (0.00s)
=== RUN   TestInvalidHasEqualQuotes
--- PASS: TestInvalidHasEqualQuotes (0.00s)
=== RUN   TestValidHasSpelledOutNumbers
--- PASS: TestValidHasSpelledOutNumbers (0.00s)
=== RUN   TestInvalidHasSpelledOutNumbers
--- PASS: TestInvalidHasSpelledOutNumbers (0.00s)
PASS
ok      validate_sentence       (cached)
```
