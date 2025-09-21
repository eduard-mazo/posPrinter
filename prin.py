import usb.core
import usb.util

VENDOR_ID = 0x0416
PRODUCT_ID = 0x5011

dev = usb.core.find(idVendor=VENDOR_ID, idProduct=PRODUCT_ID)
if dev is None:
    raise ValueError("Impresora POS58 no encontrada")

# En Windows NO hace falta is_kernel_driver_active
dev.set_configuration()

cfg = dev.get_active_configuration()
intf = cfg[(0, 0)]
ep = usb.util.find_descriptor(
    intf,
    custom_match=lambda e: usb.util.endpoint_direction(e.bEndpointAddress)
    == usb.util.ENDPOINT_OUT,
)


def send(data: bytes):
    ep.write(data)


send(b"Hola AMOR\n")
send(b"\x1b\x45\x01Negrita ON -- TE AMO\x1b\x45\x00\n")
send(b"\x1b\x45\x01Negrita ON -- TE AMO\x1b\x45\x00\n")
send(b"\x1b\x45\x01Negrita ON -- TE AMO\x1b\x45\x00\n")
send(b"\x1b\x45\x01Negrita ON -- TE AMO\x1b\x45\x00\n")
send(b"\x1b\x45\x01Negrita ON -- TE AMO\x1b\x45\x00\n")
send(b"\x1d\x56\x00")
