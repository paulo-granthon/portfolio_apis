FROM postgres:16.2

RUN apt-get update && \
    apt-get install -y \
        build-essential git postgresql-server-dev-16 python3 && \
    apt-get clean && rm -rf /var/lib/apt/lists/* # Clean up the package list

RUN git clone -b master --depth=1 https://github.com/petere/pguint.git /pguint

WORKDIR /pguint

RUN make && make install && rm -rf /pguint
# RUN apt-get purge -y build-essential git postgresql-server-dev-16 python3 && \
#     apt-get autoremove -y && apt-get clean all -y
