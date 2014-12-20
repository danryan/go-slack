# go-slack

Slack API client in Go.

Tested on all Go versions 1.1 and higher.

## Getting started

```go
package main

import "fmt"
import "github.com/danryan/go-slack/slack"

func main() {
  team := "SLACK_TEAM"
  apiKey := "SLACK_API_KEY"
  client := slack.New(team, apiKey)

  channels := client.Channels.List()

  fmt.Printf("There are %v channels.", len(channels))
}
```

## Resources

* [API documentation](http://godoc.org/github.com/danryan/go-slack/slack)
* [Bugs, questions, and feature requests](https://github.com/danryan/go-slack/issues)

## Is it any good?

[Possibly.](http://news.ycombinator.com/item?id=3067434)

## License

This library is distributed under the MIT License, a copy of which can be found in the [LICENSE](LICENSE) file.
