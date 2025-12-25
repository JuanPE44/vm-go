# Virtual Machine – Instruction Set (Opcode Table)

**Notation**
- `[]` → Operand Stack (values consumed / produced)
- `< >` → Immediate operand in bytecode (comes from the program stream)

---

## Stack & Arithmetic

| Opcode | Stack Before | Stack After | Description |
|------|-------------|-------------|-------------|
| `OP_PUSH <int>` | — | `[int]` | Push immediate integer onto the stack |
| `OP_POP` | `[int]` | — | Remove top value from stack |
| `OP_ADD` | `[int] [int]` | `[int]` | Pop two values, push sum |
| `OP_SUB` | `[int] [int]` | `[int]` | Pop two values, push subtraction |
| `OP_MUL` | `[int] [int]` | `[int]` | Pop two values, push multiplication |
| `OP_DIV` | `[int] [int]` | `[int]` | Pop two values, push division |
| `OP_PRINT` | `[int]` | — | Print top of stack (or full stack, VM-defined) |

---

## Comparisons (Boolean result: `0 = false`, `1 = true`)

| Opcode | Stack Before | Stack After | Description |
|------|-------------|-------------|-------------|
| `OP_EQ` | `[int] [int]` | `[bool]` | Equal |
| `OP_NEQ` | `[int] [int]` | `[bool]` | Not equal |
| `OP_GT` | `[int] [int]` | `[bool]` | Greater than |
| `OP_GE` | `[int] [int]` | `[bool]` | Greater or equal |
| `OP_LT` | `[int] [int]` | `[bool]` | Less than |
| `OP_LE` | `[int] [int]` | `[bool]` | Less or equal |

---

## Control Flow

| Opcode | Stack Before | Stack After | Description |
|------|-------------|-------------|-------------|
| `OP_JUMP <target>` | — | — | Unconditional jump |
| `OP_JUMP_IF_TRUE <target>` | `[bool]` | — | Jump if condition is true |
| `OP_JUMP_IF_FALSE <target>` | `[bool]` | — | Jump if condition is false |
| `OP_HALT` | — | — | Stop program execution |

---

## Memory (Variables)

| Opcode | Stack Before | Stack After | Description |
|------|-------------|-------------|-------------|
| `OP_LOAD <slot>` | — | `[int]` | Load value from memory slot onto stack |
| `OP_STORE <slot>` | `[int]` | — | Store value from stack into memory slot |

---

## Functions / Call Stack

| Opcode | Stack Before | Stack After | Description |
|------|-------------|-------------|-------------|
| `OP_CALL <target>` | — | — | Call function at target address |
| `OP_RET` | — | — | Return to caller (restore PC from call stack) |

---

## Notes

- All control-flow instructions modify the program counter (`PC`) directly.
- Boolean values are represented as integers: `0 = false`, `1 = true`.
- Stack underflow or invalid jumps should raise VM errors.
- Function code is placed outside the main execution flow and reached only via `OP_CALL`.
