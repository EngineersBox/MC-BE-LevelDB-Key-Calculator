# Sub Chunk Data Format

```
<random bytes probably for chunk info>
<serial block data>
```

## Block Data Format

Format Keys:
* `C` : characters
* `N` : number
* `B` : non-data bytes
* `*` : anything

```
1 bytes | N | format (compound)      | 0A
1 bytes | N | length                 | 00
1 bytes | B | separator              | 00
1 bytes | N | format (string)        | 08
1 bytes | N | tag name length        | 04
1 bytes | B | separator              | 00
4 bytes | C | name                   | 6E 61 6D 65
2 bytes | N | name length (short)    | XX XX
n bytes | C | block name             | <namespace>:<block>
1 bytes | N | format (compound)      | 0A
1 bytes | N | states name length     | 06
3 bytes | B | separator              | 00
6 bytes | C | states                 | 73 74 61 74 65 73
n bytes | B | states list            | <states elements>
1 bytes | B | separator/end          | 00
1 bytes | N | format (int)           | 03
1 bytes | N | version name length    | 07
3 bytes | B | separator              | 00
7 bytes | C | version                | 76 65 72 73 69 6F 6E
4 bytes | N | version number (int32) | XX XX XX XX
7 bytes | B | separator/end          | 00
```

### Formats

#### Byte (format `01`)
```
1 bytes | N | format             | 01
1 bytes | N | name length (byte) | XX
1 bytes | B | separator          | 00
n bytes | C | name               | <name>
1 bytes | N | value (byte)       | 00 or 01
```

#### Short (format `02`)
```
1 bytes | N | format             | 02
1 bytes | N | name length (byte) | XX
1 bytes | B | separator          | 00
n bytes | C | name               | <name>
2 bytes | N | value (short16)    | XX XX
```

#### Int (format `03`)
```
1 bytes | N | format             | 03
1 bytes | N | name length (byte) | XX
1 bytes | B | separator          | 00
n bytes | C | name               | <name>
4 bytes | N | value (int32)      | XX XX XX XX
```

#### Long (format `04`)
```
1 bytes | N | format             | 04
1 bytes | N | name length (byte) | XX
1 bytes | B | separator          | 00
n bytes | C | name               | <name>
8 bytes | N | value (long64)     | XX XX XX XX XX XX XX XX
```

#### Float (format `05`)
```
1 bytes | N | format             | 05
1 bytes | N | name length (byte) | XX
1 bytes | B | separator          | 00
n bytes | C | name               | <name>
4 bytes | N | value (float32)    | XX XX XX XX
```

#### Double (format `06`)
```
1 bytes | N | format             | 06
1 bytes | N | name length (byte) | XX
1 bytes | B | separator          | 00
n bytes | C | name               | <name>
8 bytes | N | value (double64)   | XX XX XX XX XX XX XX XX
```

#### Byte Array (format `07`)
```
1 bytes | N | format               | 07
1 bytes | N | name length (bytes)  | XX
1 bytes | B | separator            | 00
n bytes | C | name                 | <name>
4 bytes | N | array length (int32) | XX XX XX XX
n bytes | N | value (byte)         | [XX]
```

#### String (format `08`)
```
1 bytes | N | format              | 08
1 bytes | N | name length (byte)  | XX
1 bytes | B | separator           | 00
n bytes | C | name                | <name>
1 bytes | N | value length (byte) | XX
1 bytes | B | seperator           | 00
n bytes | C | value (string)      | <value>
```

#### List (format `09`)
```
1 bytes | N | format                | 09
1 bytes | N | name length (byte)    | XX
1 bytes | B | separator             | 00
n bytes | C | name                  | <name>
4 bytes | N | list length (int32)   | XX XX XX XX
n bytes | * | values (list payload) | [<list payload>]
```

##### List Payload
```
1 bytes | N | id                   | XX
4 bytes | N | payload size (int32) | XX XX XX XX
```

#### Compound (format `0A`)
```
1 bytes | N | format        | 0A
1 bytes | N | length        | 00
1 bytes | B | separator     | 00
n bytes | * | tags          | <tags>
1 bytes | B | separator/end | 00
```

#### Int Array (format `0B`)
```
1 bytes | N | format               | 0B
1 bytes | N | name length (bytes)  | XX
1 bytes | B | separator            | 00
n bytes | C | name                 | <name>
4 bytes | N | array length (int32) | XX XX XX XX
n bytes | N | value (int32)        | [XX XX XX XX]
```

#### Long Array (format `0C`)
```
1 bytes | N | format               | 0B
1 bytes | N | name length (bytes)  | XX
1 bytes | B | separator            | 00
n bytes | C | name                 | <name>
4 bytes | N | array length (int32) | XX XX XX XX
n bytes | N | value (long64)       | [XX XX XX XX XX XX XX XX]
```