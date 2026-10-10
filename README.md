# Fast Tools

We often need specific utilities that may be available across various websites. However, these tools are not fully integrated with our computers, and accessing them can be time-consuming. Additionally, some websites may have limited functionality.

**Fast Tools** addresses these issues by providing three main features:

1. Password Generator
2. File Converter
3. Automatic SSH Login

## Dependencies

The video and text conversion modules require **FFmpeg** to be installed and available in your terminal's `PATH`.

## Installation

You can install Fast Tools using the system package available for **Ubuntu, Windows, and macOS**.

After downloading the package, it is recommended to add the executable to your system's `PATH` so that it can be accessed from any terminal location.

## How to Use

### Password Generator

To generate a password, run:

```bash
fast-tools password
```

This command generates a password using the default settings.

You can customize the generated password using the following options:

- **`--capitalize`**: Specifies the capitalization mode:
  - `both`: Allows both uppercase and lowercase letters.
  - `up`: Allows only uppercase letters.
  - `low`: Allows only lowercase letters.
- **`--noLetters`**: Excludes letters from the password.
- **`--noNumbers`**: Excludes numbers from the password.
- **`--noSimbols`**: Excludes symbols from the password.
- **`--size`**: Specifies the password length. The default is `8` characters.

### File Converter

Fast Tools supports converting audio, text, and video files.

Use the following command:

```bash
fast-tools convert "<source file>" "<destination file>"
```

Specify the source and destination file paths, including their extensions, to determine the input and output formats.

You can also use the `*.format` pattern in a filename to select all files matching a specific format.

For example, `*.mp4` can be used to select all MP4 files.

### SSH Login

Before using this feature, it is important to understand one security consideration:

**Passwords are stored without encryption.** Therefore, if you are concerned about potential security attacks, it is recommended that you avoid using this feature to store sensitive credentials.

#### Adding an SSH Connection

To add a server connection, run:

```bash
fast-tools ssh add <label> <name> <host> <password>
```

The arguments are:

- **Label**: The identifier used to reference the connection later.
- **Name**: The server's name.
- **Host**: The server's IP address.
- **Password**: The server's password.

#### Connecting to a Server

To connect to a previously configured server, run:

```bash
fast-tools ssh enter <label>
```

Replace `<label>` with the identifier you specified when adding the connection.
