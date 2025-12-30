# VM-Go: De Expresiones Matemáticas a Bytecode

Este proyecto es una implementación completa de una **Stack-based Virtual Machine (VM)** y un **Compilador de Expresiones**. El sistema es capaz de tomar una expresión humana como `(10 + 5) * 2`, traducirla a Assembly, luego a Bytecode binario y ejecutarla en un entorno virtual.

## Arquitectura del Proyecto

El proyecto se divide en cuatro capas principales:

1.  **Frontend (Compilador)**: 
    * **Lexer**: Convierte el texto en tokens (`TokenPlus`, `TokenNumber`, etc.).
    * **Parser**: Construye un Árbol de Sintaxis Abstracta (AST) respetando la precedencia de operadores.
2.  **Middleware (Generador de Código)**:
    * Recorre el AST y genera código Assembly en formato de texto plano.
3.  **Assembler**:
    * Realiza dos pasadas sobre el Assembly para resolver etiquetas (labels) y generar el Bytecode final.
4.  **Backend (Virtual Machine)**:
    * **Execution Engine**: El corazón que procesa los OpCodes.
    * **Stack**: Memoria persistente para operaciones aritméticas y flujo de control.

---

## Pipeline de Compilación

<img width="1451" height="503" alt="image" src="https://github.com/user-attachments/assets/2511a9fc-5570-4b5c-9d48-66cc04cd8adb" />

---

## Estructura de Directorios

```text
.
├── cmd/
│   └── main.go          # Punto de entrada del programa
├── pkg/
│   ├── ast/             # Definición de los nodos del árbol (AST)
│   ├── compiler/        # Lexer, Parser y Generador de Código
│   ├── assembler/       # Traductor de Assembly a Bytecode
│   └── vm/              # Motor de ejecución y lógica de la VM
├── examples/            # Programas de prueba (.asm)
└── go.mod               # Definición del módulo de Go

```

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
| `OP_PRINT` | — | — | Print top of stack |
| `OP_DUMP` | — | — | Print  full stack, VM-defined |
| `OP_DUP` |  `[int]` |  `[int] [int]` | Duplicate top value on stack |
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
