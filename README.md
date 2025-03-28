# hits

Creates hit counter badges like ![hits](https://hits.konst.fish/api/count/incr/badge.svg?url=https://github.com/konstfish/hits)

## Usage

`GET` on `/api/count/incr/badge.svg` to increase the hit counter and return a badge or `/api/count/show/badge.svg` to return a badge. The service is available publicly (for now) at https://hits.konst.fish

### URL Parameters

| Parameter  | Default       | Description                                                                                      |
| ---------- | ------------- | ------------------------------------------------------------------------------------------------ |
| `url`      | N/A, required | defines which URL the hit should be counted for                                                  |
| `title`    | `hits`        | sets the badges description text                                                                 |
| `count_bg` | `#007ec6`     | background color of the hit counter section, note `#` needs to be url encoded (e.g. `%23b48ead`) |

### Examples

`![hits](https://hits.konst.fish/api/count/show/badge.svg?url=https://github.com/konstfish/hits)`

![hits](https://hits.konst.fish/api/count/show/badge.svg?url=https://github.com/konstfish/hits)

`![hits](https://hits.konst.fish/api/count/show/badge.svg?url=https://github.com/konstfish/hits&title=views&count_bg=%23b48ead)`

![hits](https://hits.konst.fish/api/count/show/badge.svg?url=https://github.com/konstfish/hits&title=views&count_bg=%23b48ead)
