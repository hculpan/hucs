package main

import (
  "bytes"
  "regexp"
  "fmt"
  "strings"
  "errors"
  "strconv"
  )

const PUSH_OP byte = 1
const POP_OP byte = 2
const ADD_OP byte = 3
const SUB_OP byte = 4
const MULT_OP byte = 5
const DIV_OP byte = 6
const MOD_OP byte = 7
const OUTI_OP byte = 8
const OUTC_OP byte = 9

const NO_ADDRESS_MODE byte = 0
const IMMEDIATE byte = 1
const ABSOLUTE byte = 3

func AssembleLine(line string) (*bytes.Buffer, error) {
  result := new(bytes.Buffer)
  tokens := RemoveComments(regexp.MustCompile("\\s+").Split(strings.TrimSpace(line), -1))
  var err error
  if len(tokens) > 0 {
    switch strings.ToLower(tokens[0]) {
      case "push":
        err = AssemblePush(tokens, result)
/*      case "pop":
        err = AssemblePop(tokens, result)*/
      case "add":
        err = AssembleAdd(tokens, result)
      case "sub":
        err = AssembleSub(tokens, result)
      case "mult":
        err = AssembleMult(tokens, result)
      case "div":
        err = AssembleDiv(tokens, result)
      case "mod":
        err = AssembleMod(tokens, result)
      case "outi":
        err = AssembleOuti(tokens, result)
      case "outc":
        err = AssembleOutc(tokens, result)
      default:
        return nil, fmt.Errorf("unrecognized token '%s'", tokens[0])
    }
  }

  if err != nil {
    return nil, err
  }

  return result, nil
}

// AssemblePush assembles the Push instruction
func AssemblePush(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) == 1 {
    return errors.New("Requires operand")
  } else if len(tokens) > 2 {
    return fmt.Errorf("EOL expected; %s found", tokens[2])
  }

  buf.WriteByte(PUSH_OP)

  var paramStr string = tokens[1]
  if strings.HasPrefix(paramStr, "#") {
    buf.WriteByte(IMMEDIATE)
    paramStr = paramStr[1:]
  }

  if num, err := strconv.Atoi(paramStr); err != nil {
    return fmt.Errorf("%s is not a valid value", paramStr)
  } else if num > 65535 {
    return fmt.Errorf("%d is greater than parameter capacity", num)
  } else if num > 255 {
    buf.WriteByte(byte(num % 256))
    buf.WriteByte(byte(num / 256))
  } else {
    buf.WriteByte(byte(num))
    buf.WriteByte(0)
  }

  return nil;
}

// AssembleAdd assembles the Add instruction
func AssembleAdd(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) > 1 {
    return fmt.Errorf("EOL expected; %s found", tokens[1])
  }

  buf.WriteByte(ADD_OP)

  return nil;
}

// AssembleSub assembles the Add instruction
func AssembleSub(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) > 1 {
    return fmt.Errorf("EOL expected; %s found", tokens[1])
  }

  buf.WriteByte(SUB_OP)

  return nil;
}

// AssembleMult assembles the Add instruction
func AssembleMult(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) > 1 {
    return fmt.Errorf("EOL expected; %s found", tokens[1])
  }

  buf.WriteByte(MULT_OP)

  return nil;
}

// AssembleDiv assembles the Add instruction
func AssembleDiv(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) > 1 {
    return fmt.Errorf("EOL expected; %s found", tokens[1])
  }

  buf.WriteByte(DIV_OP)

  return nil;
}

// AssembleMod assembles the Add instruction
func AssembleMod(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) > 1 {
    return fmt.Errorf("EOL expected; %s found", tokens[1])
  }

  buf.WriteByte(MOD_OP)

  return nil;
}

// AssembleOuti assembles the Add instruction
func AssembleOuti(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) > 1 {
    return fmt.Errorf("EOL expected; %s found", tokens[1])
  }

  buf.WriteByte(OUTI_OP)

  return nil;
}

// AssembleOutc assembles the Add instruction
func AssembleOutc(tokens []string, buf *bytes.Buffer) error {
  if len(tokens) > 1 {
    return fmt.Errorf("EOL expected; %s found", tokens[1])
  }

  buf.WriteByte(OUTC_OP)

  return nil;
}

// RemoveComments searches the list of tokens for
// the first item with the double-slash comment
// indicator and returns only the tokens before it
func RemoveComments(tokens []string) []string {
  for i, token := range tokens {
    if strings.HasPrefix(token, "//") {
      return tokens[:i]
    }
  }

  return tokens
}
