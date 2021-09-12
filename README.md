# cloudbuildgraph

`cloudbuildgraph` visualizes your Google Cloud Build pipelines.

fork from [patrickhoefler/cloudbuildgraph: Visualize your Google Cloud Build pipelines](https://github.com/patrickhoefler/cloudbuildgraph)
## Example Output

![Example graph](example/cloudbuild.png)

## Getting Started

### Prerequisites

- [Google Cloud Build config](https://cloud.google.com/cloud-build/docs/build-config) file in your current working directory

### Installation and Usage

Running `cloudbuildgraph` will create a `cloudbuild.pdf(png/jpeg/svg)` file in your current working directory that contains a visual graph representation of your Cloud Build pipeline.

Install

```shell
go install github.com/lirlia/cloudbuildgraph@latest
```

Usage 

```shell
  -config string
        cloudbuild config name (default "cloudbuild.yaml")
  -type string
        output type (jpeg/pdf/png/svg) (default "pdf")
```
#### Build from Source

```shell
go build
./cloudbuildgraph
```

## License

[MIT](https://github.com/lirlia/cloudbuildgraph/blob/main/LICENSE)
