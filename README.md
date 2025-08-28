# [`wit-motion` module](https://github.com/viam-modules/wit-motion)

This [wit-motion module](https://app.viam.com/module/viam/wit-motion) implements a [WitMotion](https://www.wit-motion.com/) imu movement sensor, including the [HWT905-TTL IMU](https://www.wit-motion.com/proztgjd/39.html) model using the [`rdk:component:movement_sensor` API](https://docs.viam.com/appendix/apis/components/movement_sensor/).

The `imu-wit` movement sensor model supports the following IMUs manufactured by :
- BWT61CL
- BWT901CL
- HWT901B-TTL

The `imu-wit-hwt905` movement sensor model supports the [HWT905-TTL IMU](https://www.wit-motion.com/proztgjd/39.html) manufactured by WitMotion.

Other WitMotion IMUs that communicate over serial may also work with this model but have not been tested.

> [!NOTE]
> Before configuring your movement_sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add movement_sensor / wit-motion:imu-wit to your machine](https://docs.viam.com/configure/#components).

## Configure your imu-wit movement_sensor

To configure a wit-motion movement sensor, you must set the serial path. To find your serial device path, first connect the serial device to your machine:

- On Linux, run `find /dev/serial/by-path/` to show connected serial devices, or look for your device in the output of `sudo dmesg | grep tty`. Use the device including the path to the device. Example: `"/dev/serial/by-path/usb-0:1.1:1.0"`.
- On macOS, run `ls /dev/tty* | grep -i usb` to show connected USB serial devices, `ls /dev/tty*` to browse all devices, or look for your device in the output of `sudo dmesg | grep tty`. Example: `"/dev/ttyS0"`.

```json
{
    "serial_path": "<your-port>",
    "serial_baud_rate": <int>
}
```

### Attributes

The following attributes are available for `viam:wit-motion:imu-wit` movement_sensors:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `serial_path` | string | **Required** | The full filesystem path to the serial device, starting with `/dev/ |
| `serial_baud_rate` | int | Optional | The rate at which data is sent from the sensor over the serial connection. Valid rates are `9600` and `115200`. The default rate will work for all models. _Only the HWT901B can have a different serial baud rate._ Refer to your model's data sheet. Default: `115200` |

### Example Configuration

#### `viam:wit-motion:imu-wit`

```json
  {
    "serial_path": "/dev/serial/by-path/usb-0:1.1:1.0",
    "serial_baud_rate": 115200
  }
```

#### `viam:wit-motion:imu-wit-hwt905`

```json
  {
    "serial_path": "/dev/serial/by-path/usb-0:1.1:1.0",
    "serial_baud_rate": 115200
  }
```

### Next Steps

- To test your movement_sensor, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your movement_sensor, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a movement_sensor component, explore [these tutorials](https://docs.viam.com/tutorials/).
