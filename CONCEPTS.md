# Core Concepts

## Machine
A machine is a system that executes instructions over a mutable state according to a well-defined execution model. It consists of a set of operations (instructions), a state that can be read and modified, and a control mechanism that determines the order in which instructions are executed. A machine operates by repeatedly fetching an instruction, interpreting it, and applying its effect to the current state.

In computing, a physical CPU is an example of a machine, where registers, memory, and flags form the state, and the instruction set defines the allowed operations. More abstractly, a machine can be described independently of its physical implementation, focusing only on its behavior and execution semantics.

## Virtual Machine
A virtual machine (VM) is a machine defined and implemented entirely in software that executes programs according to a specified abstract execution model. Instead of running directly on physical hardware, a virtual machine operates over a simulated state and interprets or executes a custom instruction set, often referred to as bytecode. The VM defines how instructions affect its state, independently of the underlying platform.

Virtual machines are commonly used to provide portability, isolation, and a well-defined runtime environment. Examples include language runtimes such as the Java Virtual Machine (JVM) or Lua VM, as well as custom virtual machines designed to execute domain-specific languages or experimental instruction sets.

## Instruction Set (ISA)
An Instruction Set Architecture (ISA) is the formal specification of the set of instructions that a machine or virtual machine can execute, along with their semantics and encoding. It defines which operations are available, how instructions are represented, what operands they accept, and how each instruction affects the machine state. The ISA acts as a contract between the code being executed and the execution engine.

In the context of a virtual machine, the ISA describes the bytecode format, the meaning of each opcode, and the expected behavior of instructions such as arithmetic operations, memory access, control flow, and function calls. A well-defined ISA allows programs to be generated, analyzed, and executed consistently, regardless of the internal implementation of the VM.

## Opcode
An opcode (operation code) is a numeric identifier that represents a specific instruction in a machine or virtual machine’s instruction set. It is the part of an instruction that tells the execution engine what operation to perform, such as arithmetic computation, memory access, control flow, or function invocation. Each opcode has a well-defined semantics that describes how it affects the machine state.

In a virtual machine, opcodes are typically encoded as bytes or integers within the bytecode stream and may be followed by zero or more operands. During execution, the VM fetches the opcode, decodes it, and dispatches the corresponding operation according to the rules defined by the instruction set architecture (ISA).

## Bytecode
Bytecode is a low-level, platform-independent representation of a program that is executed by a virtual machine. It consists of a sequence of opcodes and their associated operands, encoded in a compact binary format. Bytecode is designed to be easy for a virtual machine to interpret or execute, while remaining abstract enough to be generated from higher-level languages.

In a typical system, bytecode is produced by a compiler or an assembler and consumed by the VM’s execution engine. It defines the exact sequence of operations to perform, the flow of control, and how data is manipulated at runtime, without being tied to any specific physical hardware.

## Program Counter (PC)
The Program Counter (PC) is a component of a machine or virtual machine that keeps track of the address or position of the next instruction to be executed. It represents the current point of execution within a program and is updated as instructions are fetched and executed. Under normal execution, the PC advances sequentially, but control-flow instructions such as jumps, calls, and returns can modify it explicitly.

In a virtual machine, the PC typically indexes into the bytecode array and is essential for implementing control flow, function calls, and loops. Correct management of the program counter ensures that instructions are executed in the intended order and that execution can move predictably between different parts of a program.

## Stack
A stack is a last-in, first-out (LIFO) data structure used by a machine or virtual machine to store temporary values during program execution. In a stack-based execution model, operands for instructions are implicitly taken from the stack and results are pushed back onto it. This allows instructions to remain simple, as they do not need to explicitly reference their operands.

In a virtual machine, the stack is commonly used to hold intermediate computation results, function arguments, and return values. The correct management of the stack is critical, as operations such as pushing, popping, and accessing values must follow strict rules to avoid errors like stack underflow or corruption of execution state.

## Call Stack
The call stack is a stack-based structure that tracks active function or procedure calls during program execution. Each time a function is invoked, a new entry—commonly called a stack frame—is pushed onto the call stack, representing the execution context of that function. When the function returns, its frame is removed, and execution resumes at the caller’s context.

In a virtual machine, the call stack is essential for managing control flow across function calls, preserving return addresses, local variables, and arguments. It enables nested and recursive function calls by ensuring that each invocation has its own isolated execution context.

## Stack Frame
A stack frame is a data structure that represents the execution context of a single function or procedure call. It typically contains the information needed to resume execution after the function returns, such as the return address, function arguments, local variables, and any temporary values required during execution. Each active function call has its own stack frame.

In a virtual machine, stack frames are managed by the call stack and are created when a function is invoked and destroyed when it returns. Stack frames provide isolation between function calls, ensuring that local state does not interfere with other parts of the program and enabling features such as nested calls and recursion.

## ALU
An ALU (Arithmetic Logic Unit) is a component of a machine or virtual machine responsible for performing arithmetic and logical operations. These operations typically include addition, subtraction, multiplication, division, and logical comparisons such as equality, greater-than, or boolean conjunctions. The ALU operates on input values provided by the machine’s state, such as registers or a stack, and produces results that are written back to that state.

In a virtual machine, the ALU is often implemented as a logical module or set of functions rather than a physical unit. It defines the semantics of arithmetic and logical opcodes and ensures that these operations behave consistently according to the rules of the instruction set architecture (ISA).

## Engine / Execution Engine
The execution engine is the component of a machine or virtual machine responsible for driving program execution. It implements the main execution loop, commonly described as the fetch–decode–execute cycle. The engine reads instructions from the program representation (such as bytecode), decodes the current opcode, and invokes the corresponding behavior that modifies the machine state.

In a virtual machine, the execution engine coordinates all core components, including the program counter, stack, call stack, and ALU. It does not define what instructions exist, but enforces how and when they are executed, ensuring that programs follow the control flow and semantics defined by the instruction set architecture.

## Runtime
The runtime is the execution environment that provides the necessary infrastructure for a program to run on a machine or virtual machine. It encompasses all components and services that are active during program execution, including the virtual machine state, memory management, call stack handling, and built-in behaviors required by the language or system. The runtime defines how programs behave while they are running, beyond just individual instruction execution.

In the context of a virtual machine, the runtime includes the execution engine, data structures such as stacks and environments, and any supporting mechanisms needed to handle function calls, errors, or built-in operations. It acts as the bridge between the compiled program (bytecode) and the underlying execution model.

## Assembler
An assembler is a tool that translates a low-level, human-readable representation of instructions into machine-executable form, typically bytecode or machine code. The input to an assembler is an assembly-like language where each instruction closely corresponds to a specific opcode in the instruction set. The assembler resolves symbolic names such as labels or function identifiers into concrete addresses or offsets used by the execution engine.

In a virtual machine context, the assembler serves as a bridge between raw bytecode and higher-level tooling. It allows developers to write programs in a readable textual form while maintaining a one-to-one relationship with the VM’s instruction set, making it especially useful for debugging, testing, and validating the behavior of the VM.

## Compiler
A compiler is a program that translates source code written in a high-level language into a lower-level representation that can be executed by a machine or virtual machine. This translation process typically involves multiple stages, such as lexical analysis, parsing, semantic analysis, and code generation. The output of a compiler may be machine code, bytecode, or another intermediate representation.

In the context of a virtual machine, the compiler’s role is to transform language constructs—such as expressions, variables, control flow, and functions—into bytecode that conforms to the virtual machine’s instruction set architecture. The compiler defines how the abstractions of the language map onto the execution model provided by the VM.

## Lexer
A lexer (lexical analyzer) is the first stage of a compiler or interpreter that processes source code and converts it into a sequence of tokens. It reads the raw input as a stream of characters and groups them into meaningful units such as identifiers, keywords, literals, operators, and symbols. The lexer is responsible for ignoring irrelevant elements like whitespace and comments.

By transforming plain text into structured tokens, the lexer simplifies the work of later stages such as the parser. In a language implementation, the lexer defines the basic vocabulary of the language and ensures that the source code conforms to its lexical rules.

## Parser
A parser is a component of a compiler or interpreter that takes a sequence of tokens produced by the lexer and organizes them into a structured representation according to the grammatical rules of the language. This structure typically reflects the syntactic relationships between language constructs, such as expressions, statements, and blocks.

The parser validates that the token sequence forms a syntactically correct program and produces an intermediate representation, most commonly an Abstract Syntax Tree (AST). By enforcing the language grammar, the parser ensures that higher-level constructs are well-formed before further semantic analysis or code generation takes place.

## Token
A token is a fundamental unit produced by the lexer that represents a meaningful element of the source code. Each token typically consists of a type, which identifies its role in the language (such as identifier, keyword, operator, or literal), and an associated value or lexeme extracted from the source text. Tokens abstract away raw characters and provide a structured input for the parser.

Tokens allow the parser to reason about the program at a higher level than individual characters. By working with tokens, the language implementation can enforce syntax rules more easily and build structured representations such as abstract syntax trees.

## AST
An Abstract Syntax Tree (AST) is a hierarchical, tree-based representation of a program that captures its syntactic structure while omitting unnecessary details such as punctuation or formatting. Each node in the AST represents a language construct—such as an expression, statement, or function—and defines the relationships between these constructs according to the language grammar.

The AST is produced by the parser and serves as a central structure for further stages of language processing, including semantic analysis, optimization, and code generation. By working with an AST, a compiler or interpreter can reason about the program’s meaning in a structured and implementation-independent way.

## Symbol Table
A symbol table is a data structure used by a compiler or interpreter to store information about identifiers defined in a program. Identifiers may include variables, functions, parameters, or other named entities. For each symbol, the table typically records metadata such as its name, scope, type, and location or index used during code generation.

The symbol table plays a central role during semantic analysis and compilation, enabling the compiler to resolve references, enforce scope rules, and map high-level names to low-level representations such as memory slots or bytecode operands. In virtual machine–based systems, symbol tables are often used at compile time to assign fixed indices that the runtime can use efficiently.

## Scope
Scope defines the region of a program in which an identifier, such as a variable or function name, is valid and can be accessed. It determines where a symbol is visible and where it can be referenced without ambiguity. Common forms of scope include global scope, function scope, and block scope, each governing how names are introduced and resolved.

In a compiler or interpreter, scope rules are enforced using symbol tables, often organized hierarchically to reflect nested scopes. Proper scope management ensures that identifiers are correctly bound to their intended definitions and prevents naming conflicts or unintended access to variables outside their valid context.

## Function
A function is a named, reusable unit of code that encapsulates a specific computation or behavior. It may accept parameters as input, execute a sequence of instructions, and optionally return a result. Functions provide abstraction and modularity, allowing complex programs to be built from smaller, well-defined pieces.

In the context of a virtual machine and compiled language, a function represents a separate execution context. When a function is called, a new stack frame is created, arguments are bound to parameters, and execution jumps to the function’s entry point. When the function returns, control is transferred back to the caller along with any return value.

## Environment
An environment is a runtime structure that associates identifiers, such as variable names or parameter names, with their corresponding values during program execution. It represents the current bindings that are active at a given point in the program and is used to resolve variable accesses according to the language’s scope rules.

In a virtual machine or interpreter, environments are often tied to stack frames or execution contexts, ensuring that each function call has its own set of variable bindings. Environments enable correct handling of local variables, parameters, and, in more advanced systems, captured variables for closures.
