<h1 align="center">Rangoware</h1>

<p  align="center">
  <img  alt="Repository size"  src="https://img.shields.io/github/repo-size/LuanSilveiraSouza/rangoware?color=252F44&style=for-the-badge">

  <a  href="https://github.com/LuanSilveiraSouza/rangoware/commits/master">
    <img  alt="GitHub last commit"  src="https://img.shields.io/github/last-commit/LuanSilveiraSouza/rangoware?color=252F44&style=for-the-badge">
  </a>

  <img  alt="License"  src="https://img.shields.io/badge/license-MIT-252F44?&style=for-the-badge">
</p>

# :pushpin: Sumary

* [Introduction](#paperclip-introduction)
* [Technologies](#computer-technologies)
* [How to Run](#rocket-how-to-run)
* [Contribution, Bugs and Issues](#bug-contribution-bugs-and-issues)
* [License](#books-license)

# :paperclip: Introduction

Rangoware is a simple Ransomware that uses AES-256-GCM encryption and is writted in Go language.

:warning: WARNING: This software is made just for study purposes.
:warning: WARNING: If you want to run it locally for tests, take care of what directories you decide to encrypt.
:warning: WARNING: The software is distributed in MIT license. Its use is free, however <strong>the author doesn't take responsibility for any illegal use of the code by 3rd parties.</strong>

# :computer: Technologies

* [Golang](https://golang.org/)
* [Avanced Encryption Standard (AES)](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)

# :rocket: How to Run

```bash
# Clone Repository
$ git clone https://github.com/LuanSilveiraSouza/rangoware.git
```

### Configure and Compile

```bash
# Go to folder
$ cd rangoware

# Generate a AES crypto key
$ go run keygen/keygen.go

# In encrypter nd decrypter, set the variables cryptoKey, contact and dir

# Compile encrypter
$ cd encrypter
$ go build

# Compile decrypter
$ cd decrypter
$ go build
```

# :bug: Contribution, Bugs and Issues

Feel free to open new issues and colaborate with others issues in [Rangoware Issues](https://github.com/LuanSilveiraSouza/rangoware/issues)


# :books: License

Released in 2020 under [MIT License](https://opensource.org/licenses/MIT)

Made with :heart: by Luan Souza.