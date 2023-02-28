import hashlib


def parse_md5(passwd, salt="secret2023"):
    m = hashlib.md5()
    m.update((salt + passwd).encode("utf8"))
    return m.hexdigest()


if __name__ == "__main__":
    print(parse_md5("123456"))
