<p align="center">
    <a href="https://www.snapwonders.com/" target="_blank">
        <img src="https://snapwonders.com/img/logo/snap-wonders-logo-big.png" width="172" alt="snapWONDERS" />
    </a>
</p>

snapWonders — Deep digital media analysis, format conversions, steganography, scrubbing and regeneration. Providing digital media solutions


# Example Go client to snapWonders API
The objective of this repository is to provide a Go client example to snapWONDERS API. This includes a step-by-step
guide how to set up your development environment, the example source code and instructions on how to test and run.

Through the example source, you will be able to:
* Upload digital media (using resumable uploading following the Tus.io protocol)
* Perform deep digital media analysis, reveal hidden metadata, hidden metadata, copyrights, steganography and private information leakage
* Display the results from the analysis (in which you can extract and make use as needed)


# Installation and setup

## Development environment
For the development environment you will need:
* Install Visual Studio Code. You can download and install from [visual code studio](https://code.visualstudio.com/download)
* Install the Go programming language if you have not done so. Simply follow the instructions from [go programming language](https://go.dev/doc/install)
* Install the plugin into Visual Studio Code called "Go". See image below for details:
<img src="https://storage.snapwonders.com/cache/1/rOPuKm-R7IDpm8rNlb73huuqCvMZW7OE.png?mark=snap-wonders-logo.png&markpos=bottom&marky=30&markalpha=30&s=739e747e27cc388873bac4c801603108" alt="Go plugin" />

## snapWONDERS API Key
You will need the snapWONDERS API Key before you can get started:
* Signup and create an account at snapWONDERS at [signup](https://snapwonders.com/sign-up). If you wish to create account via Tor or I2P then you can do so by accessing snapWONDERS via the Tor or I2P portals. For the dark web links visit [browsing safely](https://snapwonders.com/browsing-safely)
* Under your account settings, scroll to the bottom under the section "API Settings" and click the button to generate your Auth API key
* Copy this key directly into the `main.go` file under the constant `SNAPWONDERS_API_KEY`


# Running the Go example
Once everything above is setup you should be able to simply open the workspace folder with Visual Studio Code and run or debug it. Simply hit the default hot keys `F5` to start debugging or to run directly use `Ctrl+F5`.
<img src="https://storage.snapwonders.com/cache/1/mhVnv6PkU7zYJDVyxyWwtVsiSTt0MhCG.png?mark=snap-wonders-logo.png&markpos=bottom&marky=30&markalpha=30&s=8c732fa0bdce5b38c6945b31c9e885f4" alt="Visual Code IDE" />

If you wish, you can change and provide your own digital media to upload (images and/or videos) and change the `MEDIA_PATH_FILENAME` constant contained in the `main.go` file. Otherwise the sample image provided is just a photo of me that I use on my social media accounts.

If all is well, then you should see the standard output to be something like below:
<img src="https://storage.snapwonders.com/cache/1/88POwr1GGHUVHLbJiAEptWKrwguDjZxa.png?mark=snap-wonders-logo.png&markpos=bottom&marky=30&markalpha=30&s=468d104174de055126453e8c370d4266" alt="Example for standard output display for Go client to snapWONDERS API" />

Which provides similar information as per the analyse results via the snapWONDERS website:
<img src="https://storage.snapwonders.com/cache/1/wEqYS8DopFx1zqoFfAaAa12-58Eh6OCj.png?mark=snap-wonders-logo.png&markpos=bottom&marky=30&markalpha=30&s=9599795d1494b2bac7e4a2dc09a47967" alt="Results sample as showing on the snapWONDERS website" />

# Documentation
Useful documentation can be found at:
* For endpoint, swagger UI and other source code examples can be found at [snapWONDERS developers documentation](https://snapwonders.com/snapwonders-openapi-specification)
* The actual snapWONDERS OpenAPI Specification can be found at [snapWONDERS OpenAPI Specification](https://api.snapwonders.com/site/docs)
* If you're wanting the actual JSON Schema details for the purpose of auto generating source from the schema you can use [snapWONDERS OpenAPI Specification JSON Schema](https://api.snapwonders.com/site/json-schema)
* This README.md and its content is mostly duplicated at [How to create a Golang (Go) client to integrate with the snapWONDERS API](https://snapwonders.com/resources/how-to-create-a-golang-go-client-to-integrate-with-the-snapwonders-api)


# Contact

## For security concerns
If you have spotted any security concerns then please reach out via [contacting snapWONDERS](https://snapwonders.com/contact) and set the subject to **"SECURITY CONCERNS"** and provide the information about your concerns. If you wish to contact via Tor or I2P then you can do so by accessing snapWONDERS via the Tor or I2P portals. For the dark web links visit [browsing safely](https://snapwonders.com/browsing-safely)

## For FAQ and questions
It may be possible that your questions are already answered in the [FAQ](https://snapwonders.com/faq). Be sure to check out the FAQ content first. Otherwise you may reach out via [contacting snapWONDERS](https://snapwonders.com/contact). If you wish to contact via Tor or I2P then you can do so by accessing snapWONDERS via the Tor or I2P portals. For the dark web links visit [browsing safely](https://snapwonders.com/browsing-safely)

# For contacting the author
Use this link to contact the author [Kenneth Springer](https://kennethbspringer.au/)
