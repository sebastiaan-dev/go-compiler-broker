@0xd8e43c9af60789d0;

struct Buffer {}

struct Package {

}

struct Actor {

}

struct Action {
    mode @0 :Text;
    package @1 :Package;
    deps @2 :List(Action);
    actor @3 :Actor;
    ignoreFail @4 :Bool;
    testOutput @5 :Buffer;
    args @6 :List(Text);

    triggers @7 :List(Action);

    buggyInstall @8 :Bool;

    # TODO: TryCache, here or at worker?

    objDir @9 :Text;
    target @10 :Text;
    built @11 :Text;
    # TODO: is this truly data?
    actionID @12 :Data;
}