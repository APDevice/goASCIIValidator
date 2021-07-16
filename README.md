# goASCIIValidator
goASCIIValidator is a simple ASCII validator. It will take in strings from stdio and validate that all characters fall within the ASCII range (7-bit characters by default). This progam can be used to verify that urls and email addresses have not been spoofed using non-latin characters. The "validASCII" validator itself can also be used independently if so desired (see validASCII directory).

## Flags
- -e : switch to the extended extended ASCII range (aka latin-1 supplement in the Unicode Standard)
- -c : generate a more concise output consisting of 1 for any valid string, 0 for any invalid one.

## License
GNU GPLv3
