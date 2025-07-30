# go-htmx-daisyui

Playing around with golang, htmx and hyperscript and testing a layout idea.

## Demonstrates

-   individual widget loading (htmx)
-   application menu navigation via partial loading (htmx)
-   echarts SSR and dynamic load (echarts, htmx)
-   maps (leaflet, mapbox)
-   styling, components (daisyui, tailwind)
-   themes (daisyui)
-   template fragments (htmx)

## Requirements

-   [Mapbox Access Token](https://docs.mapbox.com/help/dive-deeper/access-tokens/)

## Install

```shell
# Setup .env with MAPBOX_ACCESS_TOKEN
% npm install # install tailwind, daisyui
```

## Usage

Run the servers:

```shell
% make web      # run the webserver
```

Go to <http://localhost:8888>. Then clickety click click.

## Dev

Make sure the tailwind generator thing is running:

```shell
% make tailwind # run the tailwind generator thing
```
