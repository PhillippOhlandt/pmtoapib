# Postman to Api Blueprint Converter

A simple tool to convert Postman collection exports (**v2**) to Api Blueprint documentation.

Notice: It uses the [Aglio](https://github.com/danielgtaylor/aglio) include syntax!

## Usage

Assuming that `collection.json` is your Postman collection export.

```
pmtoapib collection.json docs
```

The `docs` folder will be created with the following content

```
├── collection-name.apib
└── responses
    └── user
        ├── post-login.json
        ├── get-info.json
        └── get-posts.json
```

The `responses` folder contains JSON files with an empty JSON object, 
because those information are not available in the Postman export.

The folders inside the `responses` folder line up with the request paths.

By default, the collection name will be used as the `.apib` filename.
It can be overridden with the third argument.

```
pmtoapib collection.json docs users
```

This will generate a `users.apib` file.

## Usage with Docker

You can also use the Docker image if you don't use Mac OS and don't want to compile it on your own.

```
docker run --rm -it -v "$PWD:/opt" phillippohlandt/pmtoapib collection.json docs
```

## Options

### Exclude Requests

To exclude a whole request from the generated `.apib` file, simply put the string `pmtoapib_exclude` 
in the request description. It will also prevent the generation of the response files for that request.
