# SVG to TSX Converter

A simple Go script to convert SVG files to React components in TSX format. It reads SVG files from an input directory, converts them into TSX components, and saves them in an output directory. Each component is named using PascalCase with "Icon" appended at the end.

## Features

- Converts SVG files into reusable React components.
- Automatically names files in PascalCase with an "Icon" suffix.
- Supports batch processing of multiple SVG files.
- Simple to use with customizable input and output directories.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/pooulad/svg2tsx.git
   cd svg2tsx
   ```

2.  Make sure you have Go installed on your machine. If not, install Go from [here](https://golang.org/doc/install).
3.  Build the project:

    ```bash
    go build -o ./bin/svg2tsx main.go
    ```
    or 

    ```bash
    make build
    ```

## Usage

Run the script with the following command:

`go run main.go --input ./path/to/svg-files --output ./path/to/output-directory`

or

After `make build`

`./bin/svg2tsx --input ./path/to/svg-files --output ./path/to/output-directory`

### Flags:

-   `--input` : Path to the directory containing SVG files (default: `./input`)
-   `--output` : Path to the directory where the TSX components will be saved (default: `./output`)

### Example:

`go run main.go --input ./assets/svg --output ./components/icons`

This will convert all SVG files from the `assets/svg` directory into TSX React components and save them in the `components/icons` directory.

## How It Works

-   The script traverses through the input directory and identifies all SVG files.
-   Each SVG file is converted into a React component in TSX format with the appropriate naming convention.
-   The components are created using the standard React `SVGProps` type for easy usage in your application.

## Contributing

Feel free to fork the repository, create issues, and submit pull requests to contribute to the project!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
