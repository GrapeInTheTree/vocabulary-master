# Vocabulary Master

Vocabulary Master is a CLI application designed to help users study and manage their vocabulary. It allows users to store, retrieve, study, and export words and their meanings.

## Features

- Store new words and their meanings
- Retrieve stored words
- Study words with a spaced repetition system
- Export words to CSV files
- Update existing word entries

## Installation

1. Ensure you have Go installed on your system (version 1.18 or later recommended).
2. Clone the repository:
   ```
   git clone https://github.com/grapeinthetree/vocabulary-master.git
   ```
3. Navigate to the project directory:
   ```
   cd vocabulary-master
   ```
4. Build the application:
   ```
   go build -o vocabulary-master cmd/main.go
   ```

## Usage

Run the application with the following command:

```bash
./vocabulary-master [command] [options]
```
Available commands:

- `store` (alias: `s`): Store new words
- `retrieve` (alias: `r`): Retrieve stored words
- `study` (alias: `st`): Study words
- `export` (alias: `e`): Export words to CSV
- `update` (alias: `u`): Update existing word entries

### Storing Words

To store new words:

```bash
./vocabulary-master store
```
Follow the prompts to enter words and their meanings. Type 'exit' when finished.

### Retrieving Words

To retrieve all words:

```bash
./vocabulary-master retrieve --all
```

To retrieve a specific word:

```bash
./vocabulary-master retrieve [word]
```

### Studying Words

To study all words:

```bash
./vocabulary-master study --all
```

To study words with a minimum retry count:

```bash
./vocabulary-master study --only-retry [count]
```

### Exporting Words

To export all words:

```bash
./vocabulary-master export --all
```

To export words with a minimum retry count:

```bash
./vocabulary-master export --only-retry [count]
```

### Updating Words

To update an existing word:

```bash
./vocabulary-master update [word] [new_meaning]
```







