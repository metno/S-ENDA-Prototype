# S-ENDA-Prototype

Read S-ENDA documentation on [Read the Docs](https://readthedocs.org/): [Welcome to S-ENDA’s documentation!](https://s-enda-documentation.readthedocs.io/en/latest/)

## Development environment

### Getting started

First install dependencies needed for starting the prototype by following these [instructions](https://s-enda-documentation.readthedocs.io/en/latest/devel_environ.html#Installation).

### Run prototype

* Check out this git repository:
    ```bash
    git clone https://github.com/metno/S-ENDA-Prototype
    ```
* Enter the repository:
    ```bash
    cd S-ENDA-Prototype
    ```
* Start development environment:
    ```bash
    vagrant up
    ```
#### Test components

Access the various components by pointing your web browser to these addresses.

* PyCSW:
    ```
    http://10.20.30.10:80
    ```
* Dynamic geoassets API:
    ```
    http://10.20.30.10:8080
    ```

### Update prototype

* Go to your existing `S-ENDA-Prototype` folder:
    ```bash
    cd S-ENDA-Prototype
    ```
* Pull the latest changes from GitHub:
    ```bash
    git pull
    ```
* Tear down, and rebuild components:
    ```bash
    vagrant up
    ```

### Stop development environment

* Stop development environment:
    ```bash
    vagrant halt
    ```

### Remove development environment

* Shut down and destroy development environment:
    ```bash
    vagrant destroy
    ```

###### vim: set spell spelllang=en:
