
## Log Processor

Lightweight, configurable, extensible logging library written in Go.

Features

- Multi-level logging: Supports multiple log levels (e.g., DEBUG, INFO, WARNING, ERROR, FATAL)
- Multi-output logging: Supports logging to multiple outputs (e.g., console, file, network)
- Built-in log rotation: Automatically rotates logs based on size or time
- Buffering and compression: Buffers and compresses logs for efficient processing
- Extensible: Easily extendable with custom log processors and formatters






 

## Usage/Examples

You can write custom log processors to process logs in different ways. Here's an example of a custom log processor that sends logs to a remote server:

package main

import (
  "encoding/json"
  "log"
  "net/http"
)

type RemoteLogProcessor struct{}

func (p *RemoteLogProcessor) Process(log *log.Entry) error {
  b, err := json.Marshal(log)
  if err != nil {
    return err
  }
  req, err := http.NewRequest("POST", "(link )", bytes.NewBuffer(b))
  if err != nil {
    return err
  }
  req.Header.Set("Content-Type", "application/json")
  resp, err := (link unavailable)(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()
  return nil
}


## License

[MIT](https://choosealicense.com/licenses/mit/)

License

This project is licensed under the MIT License.

Acknowledgments

This project was inspired by the following logging libraries:

- https://github.com/gookit/slog


Contributing

Contributions are welcome! Please open a pull request with your changes.

Authors

- https://github.com/Varnit01-dev
