# [`wit-motion` module](https://github.com/viam-modules/wit-motion)

This [wit-motion module](https://app.viam.com/module/viam/wit-motion) implements a [WitMotion](https://www.wit-motion.com/) imu movement sensor, including the [HWT905-TTL IMU](https://www.wit-motion.com/proztgjd/39.html) model using the [`rdk:component:movement_sensor` API](https://docs.viam.com/appendix/apis/components/movement_sensor/).

The `imu-wit` movement sensor model supports the following IMUs manufactured by :
- BWT61CL
- BWT901CL
- HWT901B-TTL

The `imu-wit-hwt905` movement sensor model supports the [HWT905-TTL IMU](https://www.wit-motion.com/proztgjd/39.html) manufactured by WitMotion.

Other WitMotion IMUs that communicate over serial may also work with this model but have not been tested.

## Configure your imu-wit movement_sensor

> [!NOTE]
> Before configuring your movement_sensor, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add movement_sensor / wit-motion:imu-wit to your machine](https://docs.viam.com/configure/#components).

On the new component panel, copy and paste the following attribute template into your movement_sensor's attributes field:

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
| `serial_path` | string | **Required** | The full filesystem path to the serial device, starting with <file>/dev/</file>. To find your serial device path, first connect the serial device to your machine, then:<ul><li>On Linux, run <code>ls /dev/serial/by-path/\*</code> to show connected serial devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/serial/by-path/usb-0:1.1:1.0"</code>.</li><li>On macOS, run <code>ls /dev/tty\* \| grep -i usb</code> to show connected USB serial devices, <code>ls /dev/tty\*</code> to browse all devices, or look for your device in the output of <code>sudo dmesg \| grep tty</code>. Example: <code>"/dev/ttyS0"</code>.</li></ul> |
| `serial_baud_rate` | int | Optional | The rate at which data is sent from the sensor over the serial connection. Valid rates are `9600` and `115200`. The default rate will work for all models. _Only the HWT901B can have a different serial baud rate._ Refer to your model's data sheet. <br>Default: `115200` |

## Example configuration

### `viam:wit-motion:imu-wit`
```json
  {
      "name": "<your-wit-motion-imu-wit-movementsensor-name>",
      "model": "viam:wit-motion:imu-wit",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
          "serial_path": "/dev/serial/by-path/usb-0:1.1:1.0",
          "serial_baud_rate": 115200
      },
      "depends_on": []
  }
```

### `viam:wit-motion:imu-wit-hwt905`
```json
  {
      "name": "<your-wit-motion-imu-wit-hwt905-movementsensor-name>",
      "model": "viam:wit-motion:imu-wit-hwt905",
      "type": "movement_sensor",
      "namespace": "rdk",
      "attributes": {
          "serial_path": "/dev/serial/by-path/usb-0:1.1:1.0",
          "serial_baud_rate": 115200
      },
      "depends_on": []
  }
```

### Next Steps
- To test your movement_sensor, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your movement_sensor, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a movement_sensor component, explore [these tutorials](https://docs.viam.com/tutorials/).
