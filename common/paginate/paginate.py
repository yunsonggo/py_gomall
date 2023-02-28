def paginate(page=1, size=5):
    if size <= 1:
        size = 1
    if page <= 1:
        page = 1
    offset = size * (page - 1)
    return size, offset
