# 📂 pjsw - Project Switcher CLI

**pjsw** is a simple and efficient command-line tool to manage projects, making it easy to switch, organize, and retrieve project directories. Ideal for developers who want quick access to their project folders.

## 🚀 Features

- **Switch Projects**: Quickly navigate to project directories.
- **Add Projects**: Store project names and paths in a database.
- **List All Projects**: View all stored projects with paths.

## 📦 Installation

Clone the repository and build the CLI tool:

```bash
git clone https://github.com/rishabh-j-23/pjsw.git
cd pjsw
go build -o pjsw

```

For linux user
```bash
git clone https://github.com/rishabh-j-23/pjsw.git
cd pjsw
go build -o pjsw
mv pjsw /opt/pjsw/
chmod 755 /opt/pjsw
```
in ~/.bashrc file append following at end
```bash
export PATH=$PATH:/opt/pjsw/
```

## 🛠️ Usage

### 0. Bash Integration (Optional)

To enable direct project switching from your terminal, add the following to your `~/.bashrc` or `~/.zshrc` file:

```bash
source ~/.pjsw/pjsw-switch.sh
```

After adding this, you can use the `pjsw-project` command:

```bash
pjsw-project
```

This command will launch a menu with all the projects you have added, allowing you to switch between them.


### 1. Show Help

```bash
pjsw help
```

Prints the list of available commands.

### 2. Switch to Project Directory

```bash
pjsw sw <name>
```

Copies the `cd <projectpath>` command to your clipboard for quick navigation. Paste the copied command in terminal or whichever thing you are using it for.

### 3. Add Project

```bash
pjsw add <name> <path>
```

Adds a project name and path to the database. Use `.` for the current directory.

### 4. Remove Project

```bash
pjsw rm <name>
```

Removes project from the database.

### 5. List All Projects

```bash
pjsw getall
```

Displays all added projects with their names and paths.

## 🤝 Contributing

Contributions are welcome! Feel free to fork the repository, make changes, and submit a pull request.
