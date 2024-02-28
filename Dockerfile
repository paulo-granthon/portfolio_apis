FROM postgres:16.2

### PostgreSQL container with the pguint extension

### Update the package list and install the extension build dependencies
RUN apt-get update && \
    apt-get install -y \
        build-essential git postgresql-server-dev-16 python3 && \
    apt-get clean && rm -rf /var/lib/apt/lists/* # Clean up the package list

### Clone the extension source code from GitHub
RUN git clone -b master --depth=1 https://github.com/petere/pguint.git /pguint

### Set the working directory to the extension source code
WORKDIR /pguint

### Build and install the extension, then remove the source code
RUN make && make install && rm -rf /pguint

### Clean up build dependencies of the extension
# RUN apt-get purge -y build-essential git postgresql-server-dev-16 python3 && \
#     apt-get autoremove -y && apt-get clean all -y
