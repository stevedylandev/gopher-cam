# REAL LIVE GOPHERS

![GOPHERS](https://files.stevedylan.dev/gopher-cam-demo.jpg)

A simple Go app using opencv and gioui to create a window with your active camera. 

## Install 

1. Make sure `opencv@4` is installed. 

Check your package manager for how this might be packaged; for me on Alpine it was `opencv-dev` 

2. Clone and install

```bash
git clone https://github.com/stevedylandev/gopher-cam
cd gopher-cam
go install .
```

3. Run 

```bash
gopher-cam # optional: use & after command to make it a background process
```

## License

[GNU GPL v3.0](LICENSE)
