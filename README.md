# Add-License

Command Line Program to add license file to directory

## Installation

### Option 1: Download a prebuilt binary

Go to the [Releases](https://github.com/Th3-S1lenc3/Add-License/releases) page and download the latest release for your operating system.

#### Optional
Verify your download by generating a SHA256 sum of the downloaded file and then check that matches the sum for your file on the releases page.

On *nix systems:
```
$ sha256sum <path_to_downloaded_file>
```

On Windows:
```
Powershell
$ Get-FileHash -Path <path_to_downloaded_file>

Command Prompt
$ certutil -hashfile <path_to_downloaded_file> sha256
```

### Option 2 : Clone from repository
**IMPORTANT NOTE:** cloning the master branch of the repository will download the latest unreleased, development version and may not be stable.

Download Master:
```
$ git clone https://github.com/Th3-S1lenc3/Add-License.git
$ cd Add-License
$ go build -o build/ .
```

Download stable branch:
```
$ git clone --depth 1 --branch <release version> https://github.com/Th3-S1lenc3/Add-License.git
$ cd Add-License
$ go build -o build/ .
```

## Usage

For initial usage is is recommend to run `add-license -i` to download the licenses. These will be downloaded to the app directory in your configuration directory.

To see supported licenses run `add-license -list`.

Example:
```
$ add-license -list
GNU General Public License v3.0:
 <license_description>
 To Use:
   add-license -l="gpl-3.0"
   add-license -l="GNU General Public License v3.0"
```
***License Description Omitted***

To add a license to a directory run `add-license -l="<license>"`, to specify another directory instead of the current working directory use the `-o` flag with the desired directory

Example:
```
$ add-license -l="gpl-3.0" -o ~/git/MyNewProject/
```


## License

License Data from [github/choosealicense.com](https://github.com/github/choosealicense.com)

The content of the licenses themselves (licenses/*.license) are licensed under the [Creative Commons Attribution 3.0 Unported License](https://creativecommons.org/licenses/by/3.0/), and all other content is licensed under the [GNU GPLv3 License](https://github.com/Th3-S1lenc3/Add-License/LICENSE.md).

From original license data, YAML Frontmatter was modified from Jakyll to JSON, data content was extracted verbatim and placed in corresponding [.license files](https://github.com/Th3-S1lenc3/Add-License/tree/master/licenses)
