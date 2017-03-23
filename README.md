# Postman to Api Blueprint Converter

A simple tool to convert Postman collection exports (**v2**) to Api Blueprint documentation.

Notice: It uses the [Aglio](https://github.com/danielgtaylor/aglio) include syntax!

## Usage

Assuming that `collection.json` is your Postman collection export.

```
pmtoapib -collection collection.json -destination docs
```

The `docs` folder will be created with the following content

```
├── collection-name.apib
└── responses
    ├── get-200-me.json
    ├── get-200-users.json
    ├── post-200-me.json
    └── users
        ├── get-200_userID_2-2.json
        └── get-404_userID_1-2.json
```

The `responses` folder contains JSON files with the (pretty printed) response bodies.
If a request has no exported responses (e.g. you didn't save any in Postman), 
the response file will contain an empty JSON object.

The folders inside the `responses` folder line up with the request paths and
the files are named with the following pattern: `{method}-{response name}-{last element in url}.json`.

By default, the collection name will be used as the `.apib` filename.
It can be overridden with the `-apibname` parameter.

```
pmtoapib -collection collection.json -destination docs -apibname users
```

This will generate a `users.apib` file.

## Usage with Docker

You can also use the Docker image if you don't use Mac OS and don't want to compile it on your own.

```
docker run --rm -it -v "$PWD:/opt" phillippohlandt/pmtoapib -collection collection.json -destination docs
```

## Command Line Flags

| Command | Short Version | Type | Default Value | Description |
|---------|---------------|------|---------------|-------------|
| -collection | -c | string | | Path to the Postman collection export |
| -destination | -d | string | `./` | Destination folder path for the generated files |
| -apibname | | string | | Set a custom name for the generated .apib file |
| -force-apib | | boolean | `false` | Override existing .apib files |
| -force-responses | | boolean | `false` | Override existing response files |
| -dump-request | | string | | Output the markup for a single request (Takes a request name) |


## Options

### Exclude Requests

To exclude a whole request from the generated `.apib` file, simply put the string `pmtoapib_exclude` 
in the request description. It will also prevent the generation of the response files for that request.
