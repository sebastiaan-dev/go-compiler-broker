using Go = import "/go.capnp";
@0xa3f161972afd8244;
$Go.package("main");
$Go.import("serialization/capnp");

struct Book {
    title @0 :Text;
    author @1 :Text;
}