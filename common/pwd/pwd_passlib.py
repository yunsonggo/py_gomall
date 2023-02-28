from passlib.hash import pbkdf2_sha256, pbkdf2_sha512


def parse_sha256(passwd):
    return pbkdf2_sha256.hash(passwd)


def verify_sha256(passwd, hashed):
    return pbkdf2_sha256.verify(passwd, hashed)


def parse_sha512(passwd):
    return pbkdf2_sha512.hash(passwd)


def verify_sha512(passwd, hashed):
    return pbkdf2_sha512.verify(passwd, hashed)
