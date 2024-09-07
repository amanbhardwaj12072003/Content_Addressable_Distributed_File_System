# Content Addressable Distributed File System 
A **Content-Addressable Distributed File Storage System (CADFS)** is a content-addressable, peer-to-peer (P2P) distributed storage solution built from scratch. This project provides decentralized file storage and retrieval by leveraging custom-built TCP protocols and secure peer-to-peer communication.

## Overview
A Content-Addressable Distributed File System (CADFS) is a type of decentralized storage network where files or data blocks are stored, located, and retrieved based on their content rather than traditional methods like filenames or file paths. In a CADFS, each piece of data is assigned a unique identifier derived from a cryptographic hash of its content. This unique hash serves as both the address and the key for the data, allowing the system to locate and verify the data efficiently and securely. The use of content addressing ensures that identical data is stored only once across the network, reducing redundancy and storage requirements.

In a CADFS, data is distributed across multiple nodes in a peer-to-peer network. Each node can store, replicate, and serve parts of the data, contributing to a resilient, fault-tolerant, and scalable storage solution. The absence of a central authority or server means that data can be retrieved from multiple nodes simultaneously, leveraging parallelism to enhance speed and reliability. The content-addressable model also inherently supports data integrity, as any change in the data will produce a different hash, immediately signaling any corruption or tampering. This makes the system particularly well-suited for applications where security, data integrity, and verification are critical.

Additionally, because each unique version of a file or data block is assigned a different content hash, a CADFS naturally supports versioning and immutable data storage. Changes to a file result in a new hash and thus a new entry in the system, allowing different versions to coexist without conflict or confusion. This also allows for efficient data deduplication, as identical content only needs to be stored once, even if requested or referenced by multiple users or applications.

CADFS is often used in scenarios requiring high availability, security, and decentralization, such as in blockchain technologies, peer-to-peer networks, decentralized applications (dApps), and large-scale distributed backup or archiving systems. It provides a way to store and retrieve data in a manner that is independent of any single server or organization, enhancing both privacy and resilience against data loss or censorship. Popular examples include the InterPlanetary File System (IPFS), which enables distributed file sharing across a global network, and the Git version control system, which uses content-addressable storage for managing code repositories.


## Key Characteristics

#### 1. Content Addressing:
- In a content-addressable system, each piece of data (such as a file or block of data) is identified by a unique address derived from its content. This address is usually a cryptographic hash (e.g., SHA-256) of the data itself.
- When you want to store data in a CADFS, you compute its hash and use that hash as the key to store it. When you want to retrieve the data, you use the hash to look it up.

#### 2. Decentralization:
- CADFS is often implemented as a distributed network where data is spread across multiple nodes. There is no central authority or server storing all the data; instead, each node can store parts of the data.
- Nodes in the network collaborate to store, replicate, and serve the data, making the system more resilient and scalable.

#### 3. Data Deduplication:
- Since data is addressed by its content, identical data is stored only once, even if multiple users or processes want to store the same content. This is known as deduplication, which reduces storage requirements and increases efficiency.

#### 4. Integrity and Security:
- The use of cryptographic hashes ensures that the data’s integrity is maintained. If any part of the data changes, the hash changes, allowing the system to detect tampering or corruption.
- Content addressing also facilitates secure data sharing and verification, as users can independently compute the hash of the data to confirm its authenticity.

#### 5. Versioning:
- Since the content address is tied to the specific state of the data, any modification results in a new hash. This naturally provides versioning, where different versions of the same file are stored separately and are retrievable by their unique content hashes.

#### 6. Efficient Data Distribution and Retrieval:
- In a CADFS, data can be easily replicated and distributed across multiple nodes. The retrieval process can leverage parallelism by fetching parts of the data from multiple sources simultaneously, reducing latency and improving fault tolerance.

## Features

#### 1. Content Addressability:
Files are identified by the hash of their content, ensuring data integrity and efficient lookups.

#### 2. Custom TCP Protocol: 
Implements a TCP transport layer for communication between nodes, enabling a secure and robust network.

#### 3. Decentralized Storage:
Supports P2P communication for distributing files across multiple nodes, reducing central points of failure.

#### 4. Data Security:
Utilizes encryption for secure data transmission and storage.

#### 5. Dynamic Peer Management: 
Facilitates dynamic addition and removal of peers to maintain network scalability and resilience.

#### 6. File Replication and Retrieval:
Provides efficient file replication and retrieval across nodes in the network using a content-based addressing mechanism.

#### 7. Custom Message Encoding:
Uses Gob encoding for efficient serialization and deserialization of messages.


## Folder Structure

```
Distributed_File_Storage/
│
├── :3000_network/
│   ├── ID/
│        ├── file_1
│        ├── file_2
│        ├── file_3
├── :3001_network/
│   ├── ID/
│        ├── file_1
│        ├── file_2
│        ├── file_3
├── :3002_network/
│   ├── ID/
│        ├── file_1
│        ├── file_2
│        ├── file_3
├── bin/
│   ├── dfs 
├── p2p/
│   ├── encoding.go
│   ├── handshake.go
│   ├── message.go
│   ├── tcp_transport.go
│   ├── tcp_transport_test.go
│   └── transport.go
│
├── crypto.go
├── crypto_test.go
├── store.go
├── store_test.go
├── server.go
├── main.go
├── go.mod
├── go.sum 
├── Makefile
├── gitignore
└── README.md
 
```
