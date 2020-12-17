package EncryptedFileSharingSystem
import (
        "github.com/cs161-staff/userlib"
        "encoding/json"
        "encoding/hex"
        "github.com/google/uuid"
        "strings"
        "errors"
        _ "strconv"
)
// Helper function: Takes the first 16 bytes and
// converts it into the UUID type
func bytesToUUID(data []byte) (ret uuid.UUID) {
        for x := range ret {
                ret[x] = data[x]
        }
        return
}

// The structure definition for a user record
type User struct {
        Username string
        Password string
        ID userlib.UUID
        MACID userlib.UUID
        Pub userlib.PKEEncKey
        Priv userlib.PKEDecKey
        SignPub userlib.DSVerifyKey
        SignPriv userlib.DSSignKey
                Files_uuid userlib.UUID
                Files_symkey []byte
                Files_macuuid userlib.UUID
                Files_mackey []byte
}
type Files struct {
    MyFiles map[string]userlib.UUID
    FilesSymmetric map[userlib.UUID][]byte
    Filename_key []byte
}
type Root struct{
                Start_id userlib.UUID
                Start_macid userlib.UUID
                Start_symkey []byte
                Start_mackey []byte
        Mac_id userlib.UUID
        Owner string
        Friends_uuid map[string]userlib.UUID
                Friends_sym map[userlib.UUID][]byte
        NodeSymKey []byte
        RootMacKey []byte
}
type Node struct {
        Node_id userlib.UUID
        Mac_id userlib.UUID
        NextNodeSym []byte
        NodeMacKey []byte
        Next *Node
}
// This creates a user.  It will only be called once for a user
// (unless the keystore and datastore are cleared during testing purposes)
// It should store a copy of the userdata, suitably encrypted, in the
// datastore and should store the user's public key in the keystore.
// The datastore may corrupt or completely erase the stored
// information, but nobody outside should be able to get at the stored
// User data: the name used in the datastore should not be guessable
// without also knowing the password and username.
// You are not allowed to use any global storage other than the
// keystore and the datastore functions in the userlib library.
// You can assume the password has strong entropy, EXCEPT
// the attackers may possess a precomputed tables containing
// hashes of common passwords downloaded from the internet.
func InitUser(username string, password string) (userdataptr *User, err error) {
        var userdata User
        userdataptr = &userdata
// If the user is already in the keystore error
        _, ok1 := userlib.KeystoreGet(username)
        if ok1 {return nil, errors.New(strings.ToTitle("Init: User already here"))}
// Set user attributes
                var ok2 error
        userdata.Pub, userdata.Priv, ok2 = userlib.PKEKeyGen()

                if ok2!=nil {return nil, errors.New(strings.ToTitle("Init: error loading public keys"))}
        userlib.KeystoreSet(username + "ek", userdata.Pub)
                var ok3 error
        userdata.SignPriv, userdata.SignPub, ok3 = userlib.DSKeyGen()

                if ok3!=nil {return nil, errors.New(strings.ToTitle("Init: error loading signed keys"))}
        userlib.KeystoreSet(username + "vk", userdata.SignPub)
        userdata.Username = username
        userdata.Password = password
                userdata.Files_uuid = uuid.New()
                userdata.Files_symkey = userlib.RandomBytes(16)
                userdata.Files_macuuid = uuid.New()
                userdata.Files_mackey = userlib.RandomBytes(16)
//Set up the where the user will be storing files
                my_files := make(map[string]userlib.UUID)
                files_symkeys := make(map[userlib.UUID][]byte)
                filename_key := userlib.RandomBytes(16)
                files := &Files{MyFiles: my_files, FilesSymmetric: files_symkeys, Filename_key: filename_key}
// Encrypt the users stored files in the data store and store its mac
                iv1 := userlib.RandomBytes(16)
                files_data_bytes, ok4 := json.Marshal(files)

                if ok4!=nil {return nil, errors.New(strings.ToTitle("Init: error marshaling files"))}
                files_cipher := userlib.SymEnc(userdata.Files_symkey, iv1, files_data_bytes)
                userlib.DatastoreSet(userdata.Files_uuid, files_cipher)
                files_hmac, ok5 := userlib.HMACEval(userdata.Files_mackey, files_cipher)

                if ok5!=nil {return nil, errors.New(strings.ToTitle("Init: can not hmac eval"))}
                userlib.DatastoreSet(userdata.Files_macuuid, files_hmac)
// Set the uuid for the user in datastore
        username_bytes := []byte(username)
        salt1 := username_bytes[:3]
        salt2 := username_bytes[4:7]
        user_key := userlib.Argon2Key([]byte(password), append([]byte("PWN" + username), salt1...), 16)
        user_id, ok6 := uuid.FromBytes(user_key)

                if ok6!=nil {return nil, errors.New(strings.ToTitle("Init: error getting user uuid"))}
        userdata.ID = user_id
// Set the uuid for the mac of the user
        mac_key := userlib.Argon2Key([]byte(password), append([]byte("HMAC" + username), salt2...), 16)
        mac_id, ok7 := uuid.FromBytes(mac_key)

                if ok7!=nil {return nil, errors.New(strings.ToTitle("Init: error getting user hmac uuid"))}
        userdata.MACID = mac_id
// Encrypt the user information with Symmetric key encryption
            iv2 := userlib.RandomBytes(16)
        user_data_bytes, ok8 := json.Marshal(userdata)


                if ok8!=nil {return nil, errors.New(strings.ToTitle("Init: error marshaling userdata"))}
            user_cipher := userlib.SymEnc(user_key, iv2, user_data_bytes)
        userlib.DatastoreSet(user_id, user_cipher)
// Compute HMAC for the userdata to check for malicious activity and store in the datastore
        users_hmac, ok9 := userlib.HMACEval(mac_key, user_cipher)

                if ok9!=nil {return nil, errors.New(strings.ToTitle("Init: can not hmac eval"))}
            userlib.DatastoreSet(mac_id, users_hmac)
        return userdataptr, nil
}
// This fetches the user information from the Datastore.  It should
// fail with an error if the user/password is invalid, or if the user
// data was corrupted, or if the user can't be found.
func GetUser(username string, password string) (userdataptr *User, err error) {
        var userdata User
// Fail if users public key is not in the keystore
        _, ok1 := userlib.KeystoreGet(username + "ek")
        if !ok1 {return nil, errors.New(strings.ToTitle("Get: Users ek not in the keystore"))}
// Fail if user verify key is not in the keystore
        _, ok2 := userlib.KeystoreGet(username + "vk")
        if !ok2 {return nil, errors.New(strings.ToTitle("Get: Users vk not in the keystore"))}
// Fail if the username/password is incorrect(wrong uuid)
         username_bytes := []byte(username)
        salt1 := username_bytes[:3]
        salt2 := username_bytes[4:7]

        user_key := userlib.Argon2Key([]byte(password), append([]byte("PWN" + username), salt1...), 16)
        user_id, ok3 := uuid.FromBytes(user_key)

                if ok3!=nil {return nil, errors.New(strings.ToTitle("Get: cant get from bytes for user"))}
        ciphertext, ok4 := userlib.DatastoreGet(user_id)

                if !ok4 {return nil, errors.New(strings.ToTitle("Get: User username/password is incorrect"))}
// Check if user data has been compromised
        mac_key := userlib.Argon2Key([]byte(password), append([]byte("HMAC" + username), salt2...), 16)
        mac_id, ok5 := uuid.FromBytes(mac_key)

                if ok5!=nil {return nil, errors.New(strings.ToTitle("Get: cant get mac from bytes"))}
        users_stored_mac, ok6 := userlib.DatastoreGet(mac_id)

                if !ok6 {return nil, errors.New(strings.ToTitle("Get: No user hmac"))}
//Compare ciphertext HMAC with the stored HMAC for user
        ciphertext_mac, ok7 := userlib.HMACEval(mac_key, ciphertext)

                if ok7!=nil {return nil, errors.New(strings.ToTitle("Get: Cant eval ciphertext"))}

                if !userlib.HMACEqual(users_stored_mac, ciphertext_mac) {
                    return nil, errors.New(strings.ToTitle("Corrupted: User HMACs not equal"))
                }
// Fill the userdata with the decrypted bytes from datastore
            decrypted_user_data := userlib.SymDec(user_key, ciphertext)
        json.Unmarshal(decrypted_user_data, &userdata)
// Return a pointer to the user data
        userdataptr = &userdata
        return userdataptr, nil
}
// This stores a file in the datastore.
// The plaintext of the filename + the plaintext and length of the filename
// should NOT be revealed to the datastore!
func (userdata *User) StoreFile(filename string, data []byte) {
            // Get the users files from the data store
                var files Files
                var root *Root
                var node *Node
                var this_file_uuid userlib.UUID
                var this_file_sym []byte
                enc_files_bytes, ok1 := userlib.DatastoreGet(userdata.Files_uuid)

                if !ok1 {return}
                dec_files_bytes := userlib.SymDec(userdata.Files_symkey, enc_files_bytes)
                json.Unmarshal(dec_files_bytes, &files)
            //Check if the user is storing a completly new file
                filename_hash, ok2 := userlib.HMACEval(files.Filename_key, []byte(filename))

                if ok2!=nil {return}
                filename_hash_string := hex.EncodeToString(filename_hash)

                if _, ok3 := files.MyFiles[filename_hash_string]; !ok3 {
          // The case where the user does not have this file
                                this_file_uuid = uuid.New()
                                this_file_sym = userlib.RandomBytes(16)
                                files.MyFiles[filename_hash_string] = this_file_uuid
                                files.FilesSymmetric[this_file_uuid] = this_file_sym
                                node = &Node{Node_id: uuid.New(), Mac_id: uuid.New(), NextNodeSym: userlib.RandomBytes(16),
                                                            NodeMacKey: userlib.RandomBytes(16), Next: nil}
                                root = &Root{Start_id: uuid.New(), Start_macid: uuid.New(), Start_symkey: userlib.RandomBytes(16),
                                                            Start_mackey: userlib.RandomBytes(16), Mac_id: uuid.New(), Owner: userdata.Username,
                                                          Friends_uuid: make(map[string]userlib.UUID), Friends_sym: make(map[userlib.UUID][]byte),
                                                            NodeSymKey: userlib.RandomBytes(16), RootMacKey: userlib.RandomBytes(16)}
        } else {
            // The case where the owner or friend already has the file stored
                    this_file_uuid = files.MyFiles[filename_hash_string]
                    this_file_sym = files.FilesSymmetric[this_file_uuid]
            // Get the root of the file to store
                    encrypted_root, ok4 := userlib.DatastoreGet(this_file_uuid)

                    if !ok4 {return}
                    decrypted_root := userlib.SymDec(this_file_sym, encrypted_root)
                    json.Unmarshal(decrypted_root, &root)
            //Get the entire node structure for this file
                    encrypted_node, ok5 := userlib.DatastoreGet(root.Start_id)

                    if !ok5 {return}
                    decrypted_node := userlib.SymDec(root.Start_symkey, encrypted_node)
                    json.Unmarshal(decrypted_node, &node)
                    node.Next = nil
                }
            //Encrypt the bytes and store it in the node data store with the corresponding hmac
                iv1 := userlib.RandomBytes(16)
                encrypted_data := userlib.SymEnc(root.NodeSymKey, iv1, data)
                userlib.DatastoreSet(node.Node_id, encrypted_data)
                data_hmac, ok6 := userlib.HMACEval(node.NodeMacKey, encrypted_data)

                if ok6!=nil {return}
                userlib.DatastoreSet(node.Mac_id, data_hmac)
            //Encrypt the entrire node struct in the data store with corresponding hmac
                node_struct_bytes, ok7 := json.Marshal(node)

                if ok7!=nil {return}
                iv2 := userlib.RandomBytes(16)
                encrypted_node_struct := userlib.SymEnc(root.Start_symkey, iv2, node_struct_bytes)
                userlib.DatastoreSet(root.Start_id, encrypted_node_struct)
                node_struct_hmac, ok8 := userlib.HMACEval(root.Start_mackey, encrypted_node_struct)

                if ok8!=nil {return}
                userlib.DatastoreSet(root.Start_macid, node_struct_hmac)
            //Encrypt the root and put it in the datastore with corresponding hmac
                root_mac_key := root.RootMacKey
                root_mac_uuid := root.Mac_id
                root_bytes, ok9 := json.Marshal(root)

                if ok9!=nil {return}
                iv3 := userlib.RandomBytes(16)
                encrypted_root := userlib.SymEnc(this_file_sym, iv3, root_bytes)
                userlib.DatastoreSet(this_file_uuid, encrypted_root)
                root_struct_hmac, ok10 := userlib.HMACEval(root_mac_key, encrypted_root)

                if ok10!=nil {return}
                userlib.DatastoreSet(root_mac_uuid, root_struct_hmac)
            //Encrypt the newest version of the users files with the corresponding hmac
                files_in_bytes, ok11 := json.Marshal(files)

                if ok11!=nil {return}
                iv4 := userlib.RandomBytes(16)
                encrypted_updated_files := userlib.SymEnc(userdata.Files_symkey, iv4, files_in_bytes)
                userlib.DatastoreSet(userdata.Files_uuid, encrypted_updated_files)
                files_updated_hmac, ok12 := userlib.HMACEval(userdata.Files_mackey, encrypted_updated_files)

                if ok12!=nil {return}
                userlib.DatastoreSet(userdata.Files_macuuid, files_updated_hmac)
        return
}
// This adds on to an existing file.
//
// Append should be efficient, you shouldn't rewrite or reencrypt the
// existing file, but only whatever additional information and
// metadata you need.
func (userdata *User) AppendFile(filename string, data []byte) (err error) {
    // Get the users files from the data store
            var files Files
            enc_files_bytes, ok1 := userlib.DatastoreGet(userdata.Files_uuid)

            if !ok1 {return errors.New(strings.ToTitle("Append: no user files"))}
            dec_files_bytes := userlib.SymDec(userdata.Files_symkey, enc_files_bytes)
            json.Unmarshal(dec_files_bytes, &files)
        filename_hash, ok2 := userlib.HMACEval(files.Filename_key, []byte(filename))

            if ok2!=nil {return errors.New(strings.ToTitle("Append: cant hmac files with key"))}
        filename_hash_string := hex.EncodeToString(filename_hash)
    // Check if this file even exists and can append

            if  _, ok3 := files.MyFiles[filename_hash_string]; !ok3 {
                return errors.New(strings.ToTitle("Can't append, no uuid for this file"))
            }
            this_file_uuid := files.MyFiles[filename_hash_string]
            this_file_sym := files.FilesSymmetric[this_file_uuid]
    // Get the root of the file to store
            var root Root
            encrypted_root, ok4 := userlib.DatastoreGet(this_file_uuid)

            if !ok4 {return errors.New(strings.ToTitle("Can't append, no root detected for the file"))}
            decrypted_root := userlib.SymDec(this_file_sym, encrypted_root)
            json.Unmarshal(decrypted_root, &root)
    //Get the entire node structure for this file
            var node *Node
            encrypted_node, ok5 := userlib.DatastoreGet(root.Start_id)

            if !ok5 {return errors.New(strings.ToTitle("Can't append, no node for the file exists"))}
            decrypted_node := userlib.SymDec(root.Start_symkey, encrypted_node)
            json.Unmarshal(decrypted_node, &node)
  // Set the next pointer of the file to hold the appended bytes
          appended_node := &Node{Node_id: uuid.New(), Mac_id: uuid.New(), NextNodeSym: userlib.RandomBytes(16),
                                                 NodeMacKey: userlib.RandomBytes(16), Next: nil}
            go_node := node
        for go_node.Next != nil {
                go_node = go_node.Next
        }
        symkey_for_appendednode := go_node.NextNodeSym
        go_node.Next = appended_node
    //Encrypt the bytes and store it in the node data store with the corresponding hmac for the new appended node
            iv1 := userlib.RandomBytes(16)
            encrypted_data := userlib.SymEnc(symkey_for_appendednode, iv1, data)
            userlib.DatastoreSet(appended_node.Node_id, encrypted_data)
            data_hmac, ok6 := userlib.HMACEval(appended_node.NodeMacKey, encrypted_data)

            if ok6!=nil {return errors.New(strings.ToTitle("Append: cant eval appended node"))}
            userlib.DatastoreSet(appended_node.Mac_id, data_hmac)
    //Encrypt the entrire node struct in the data store with corresponding hmac
            node_struct_bytes, ok7 := json.Marshal(node)

            if ok7!=nil {return errors.New(strings.ToTitle("Append: cant marshal the node"))}
            iv2 := userlib.RandomBytes(16)
            encrypted_node_struct := userlib.SymEnc(root.Start_symkey, iv2, node_struct_bytes)
            userlib.DatastoreSet(root.Start_id, encrypted_node_struct)
            node_struct_hmac, ok8 := userlib.HMACEval(root.Start_mackey, encrypted_node_struct)

            if ok8!=nil {return errors.New(strings.ToTitle("Append: cant eval the root"))}
            userlib.DatastoreSet(root.Start_macid, node_struct_hmac)
  return
}
// This loads a file from the Datastore.
//
// It should give an error if the file is corrupted in any way.
func (userdata *User) LoadFile(filename string) (data []byte, err error) {
     // Get the users files from the data store
              var files Files
              enc_files_bytes, ok1 := userlib.DatastoreGet(userdata.Files_uuid)

                if !ok1 {return nil, errors.New(strings.ToTitle("Load: No Files in datastore"))}

                if len(userdata.Files_symkey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(enc_files_bytes)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
              dec_files_bytes := userlib.SymDec(userdata.Files_symkey, enc_files_bytes)
                stored_files_hmac, ok2 := userlib.DatastoreGet(userdata.Files_macuuid)

                if !ok2 {return nil, errors.New(strings.ToTitle("Load: No Files hmac in datastore"))}
                if len(userdata.Files_mackey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(enc_files_bytes)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                compute_files_hmac, ok3 := userlib.HMACEval(userdata.Files_mackey, enc_files_bytes)

                if ok3!=nil {return nil, errors.New(strings.ToTitle("Load: Cant compute files hmac"))}

                if !userlib.HMACEqual(stored_files_hmac, compute_files_hmac) {
                    return nil, errors.New(strings.ToTitle("Load: Files structure is corrupted"))}
              json.Unmarshal(dec_files_bytes, &files)
                if len(files.Filename_key)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
              filename_hash, ok4 := userlib.HMACEval(files.Filename_key, []byte(filename))

                if ok4!=nil {return nil, errors.New(strings.ToTitle("Load: cant compute files hmac"))}
              filename_hash_string := hex.EncodeToString(filename_hash)
    //Check if the file exists in the Datastore

                if  _, ok5 := files.MyFiles[filename_hash_string]; !ok5 {
                    return nil, errors.New(strings.ToTitle("Load: No root uuid can be found"))}
                this_file_uuid := files.MyFiles[filename_hash_string]

                if _, ok6 := files.FilesSymmetric[this_file_uuid]; !ok6 {
                    return nil, errors.New(strings.ToTitle("Load: No root key can be found"))
                }
                this_file_sym := files.FilesSymmetric[this_file_uuid]
    // Grab the root for the file
                var root Root
                encrypted_root, ok7 := userlib.DatastoreGet(this_file_uuid)

                if !ok7 {return nil, errors.New(strings.ToTitle("Load: There is no root"))}

                if len(this_file_sym)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(encrypted_root)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                decrypted_root := userlib.SymDec(this_file_sym, encrypted_root)
                json.Unmarshal(decrypted_root, &root)
                stored_root_hmac, ok8 := userlib.DatastoreGet(root.Mac_id)

                if !ok8 {return nil, errors.New(strings.ToTitle("Load: There is no hmac"))}

                if len(root.RootMacKey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(encrypted_root)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                compute_root_hmac, ok9 := userlib.HMACEval(root.RootMacKey, encrypted_root)

                if ok9!=nil {return nil, errors.New(strings.ToTitle("Load: Not correct hmac eval"))}

                if !userlib.HMACEqual(stored_root_hmac, compute_root_hmac) {
                    return nil, errors.New(strings.ToTitle("Load: Root structure is corrupted"))}
    //Get the entire node structure for this file
                var node *Node
                encrypted_node, ok10 := userlib.DatastoreGet(root.Start_id)

                if !ok10 {return nil, errors.New(strings.ToTitle("Load: No first node in datastore"))}

                if len(root.Start_symkey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(encrypted_node)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                decrypted_node := userlib.SymDec(root.Start_symkey, encrypted_node)
                stored_nodestruct_hmac, ok11 := userlib.DatastoreGet(root.Start_macid)

                if !ok11 {return nil, errors.New(strings.ToTitle("Load: No root start id"))}
                if len(root.Start_mackey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(encrypted_node)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                compute_nodestruct_hmac, ok12 := userlib.HMACEval(root.Start_mackey, encrypted_node)

                if ok12!=nil {return nil, errors.New(strings.ToTitle("Load: nodestruct hmac cant eval"))}

                if !userlib.HMACEqual(stored_nodestruct_hmac, compute_nodestruct_hmac) {
                    return nil, errors.New(strings.ToTitle("Load: Full Node structure is corrupted"))}
                json.Unmarshal(decrypted_node, &node)
  // Get the hmac of the bytes of the first node and compare
                enc_bytes, ok13 := userlib.DatastoreGet(node.Node_id)

                if !ok13 {return nil, errors.New(strings.ToTitle("Load: cant get first node id"))}

                if len(root.NodeSymKey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(enc_bytes)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                dec_bytes := userlib.SymDec(root.NodeSymKey, enc_bytes)
                node_hmac, ok14 := userlib.DatastoreGet(node.Mac_id)

                if !ok14 {return nil, errors.New(strings.ToTitle("Load: cant get first node mac id"))}
                if len(node.NodeMacKey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                if !(len(enc_bytes)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                compute_node_hmac, ok15 := userlib.HMACEval(node.NodeMacKey, enc_bytes)

                if ok15!=nil {return nil, errors.New(strings.ToTitle("Load: Cant eval first node"))}

                if !userlib.HMACEqual(node_hmac, compute_node_hmac) {
                    return nil, errors.New(strings.ToTitle("Load: First Node bytes is corrupted"))}
        bytes_of_file := make([]byte, 0)
                bytes_of_file = append(bytes_of_file, dec_bytes...)
  // iterate through the file pointer and check if there are more nodes and accumulate the data bytes
                go_node := node
        for go_node.Next != nil {
                enc_bytes, ok16 := userlib.DatastoreGet(go_node.Next.Node_id)
                                if !ok16 {return nil, errors.New(strings.ToTitle("Load: cant get next node"))}

                                if len(go_node.NextNodeSym)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                                if !(len(enc_bytes)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                dec_bytes = userlib.SymDec(go_node.NextNodeSym, enc_bytes)
                node_hmac, ok17 := userlib.DatastoreGet(go_node.Next.Mac_id)

                                if !ok17 {return nil, errors.New(strings.ToTitle("Load: cant get next node mac"))}
                                if len(go_node.Next.NodeMacKey)!=16 {return nil, errors.New(strings.ToTitle("Load: bad length key"))}
                                if !(len(enc_bytes)>0) {return nil, errors.New(strings.ToTitle("Load: bad bytes"))}
                compute_node_hmac, ok18 := userlib.HMACEval(go_node.Next.NodeMacKey, enc_bytes)

                                if ok18!=nil {return nil, errors.New(strings.ToTitle("Load: cant eval hmacs next node"))}

                                if !userlib.HMACEqual(node_hmac, compute_node_hmac) {
                                    return nil, errors.New(strings.ToTitle("Load: bytes while iterating file are different"))}
                bytes_of_file = append(bytes_of_file, dec_bytes...)
                go_node = go_node.Next
        }
        return bytes_of_file, nil
}
// This creates a sharing record, which is a key pointing to something
// in the datastore to share with the recipient.
// This enables the recipient to access the encrypted file as well
// for reading/appending.
// Note that neither the recipient NOR the datastore should gain any
// information about what the sender calls the file.  Only the
// recipient can access the sharing record, and only the recipient
// should be able to know the sender.
func (userdata *User) ShareFile(filename string, recipient string) (
        magic_string string, err error) {
        // Basic checks to check that the user exists
                recipients_pubkey, ok1 := userlib.KeystoreGet(recipient + "ek")
                if !ok1 {return "", errors.New(strings.ToTitle("Share: without recipients pubkey"))}
                _, ok2 := userlib.KeystoreGet(recipient + "vk")
              if !ok2 {return "", errors.New(strings.ToTitle("Share: without recipients verify key"))}
        // Get my own files from the data store
                var files Files
                enc_files_bytes, ok3 := userlib.DatastoreGet(userdata.Files_uuid)

                if !ok3 {return "", errors.New(strings.ToTitle("Share: no file struct found"))}
                dec_files_bytes := userlib.SymDec(userdata.Files_symkey, enc_files_bytes)
                stored_files_hmac, ok4 := userlib.DatastoreGet(userdata.Files_macuuid)

                if !ok4 {return "", errors.New(strings.ToTitle("Share: cant eval files"))}
                compute_files_hmac, ok5 := userlib.HMACEval(userdata.Files_mackey, enc_files_bytes)

                if ok5!=nil {return "", errors.New(strings.ToTitle("Share: cant eval the mac"))}

                if !userlib.HMACEqual(stored_files_hmac, compute_files_hmac) {
                    return "", errors.New(strings.ToTitle("Share: Files structure is corrupted"))}
                json.Unmarshal(dec_files_bytes, &files)
                filename_hash, ok6 := userlib.HMACEval(files.Filename_key, []byte(filename))

                if ok6!=nil {return "", errors.New(strings.ToTitle("Share: cant eval the files"))}
                filename_hash_string := hex.EncodeToString(filename_hash)
        //Check if the file exists in the Datastore
                if  _, ok7 := files.MyFiles[filename_hash_string]; !ok7 {
                    return "", errors.New(strings.ToTitle("Share: Nothing stored to share"))}
        // Grab the root for the file
                this_file_uuid := files.MyFiles[filename_hash_string]
                this_file_sym := files.FilesSymmetric[this_file_uuid]
                var my_root Root
                encrypted_root, ok8 := userlib.DatastoreGet(this_file_uuid)

                if !ok8 {return "", errors.New(strings.ToTitle("Share: no root in datastore"))}
                decrypted_root := userlib.SymDec(this_file_sym, encrypted_root)
                json.Unmarshal(decrypted_root, &my_root)
                stored_root_hmac, ok9 := userlib.DatastoreGet(my_root.Mac_id)

                if !ok9 {return "", errors.New(strings.ToTitle("Share: cant get root mac id"))}
                compute_root_hmac, ok10 := userlib.HMACEval(my_root.RootMacKey, encrypted_root)

                if ok10!=nil {return "", errors.New(strings.ToTitle("Share: cant eval the root mac key"))}

                if !userlib.HMACEqual(stored_root_hmac, compute_root_hmac) {
                    return "", errors.New(strings.ToTitle("Share: Root structure is corrupted"))}
    // Create the correct root for my friend
                friend_root := &Root{Start_id: my_root.Start_id, Start_macid: my_root.Start_macid, Start_symkey: my_root.Start_symkey,
                                                         Start_mackey: my_root.Start_mackey, Mac_id: uuid.New(), Owner: my_root.Owner,
                                                         Friends_uuid: make(map[string]userlib.UUID), Friends_sym: make(map[userlib.UUID][]byte),
                                                         NodeSymKey: my_root.NodeSymKey, RootMacKey: userlib.RandomBytes(16)}
    // Update my_root lists with the friends information
                bytes_for_uuid := userlib.RandomBytes(16)
                friend_root_uuid, ok11 := uuid.FromBytes(bytes_for_uuid)

                if ok11!=nil {return "", errors.New(strings.ToTitle("Share: wrong bytes from uuid"))}
            my_root.Friends_uuid[recipient] = friend_root_uuid
                friend_root_sym := userlib.RandomBytes(16)
                my_root.Friends_sym[friend_root_uuid] = friend_root_sym
        //encrypt my updated root and put in the data store with corresponding hmac
                myroot_mac_key := my_root.RootMacKey
                myroot_mac_uuid := my_root.Mac_id
                myroot_bytes, ok12 := json.Marshal(my_root)

                if ok12!=nil {return "", errors.New(strings.ToTitle("Share: cant marshal the root"))}
                iv1 := userlib.RandomBytes(16)
                encrypted_myroot := userlib.SymEnc(this_file_sym, iv1, myroot_bytes)
                userlib.DatastoreSet(this_file_uuid, encrypted_myroot)
                myroot_struct_hmac, ok13 := userlib.HMACEval(myroot_mac_key, encrypted_myroot)

                if ok13!=nil {return "", errors.New(strings.ToTitle("Share: cant eval encrypted root"))}
                userlib.DatastoreSet(myroot_mac_uuid, myroot_struct_hmac)
        // Encrypt and store my friend root for the file
            friendroot_mac_key := friend_root.RootMacKey
                friendroot_mac_uuid := friend_root.Mac_id
            friendroot_inbytes, ok14 := json.Marshal(friend_root)

                if ok14!=nil {return "", errors.New(strings.ToTitle("Share: cant marshal friends root"))}
                iv2 := userlib.RandomBytes(16)
                encrypted_friendroot := userlib.SymEnc(friend_root_sym, iv2, friendroot_inbytes)
                userlib.DatastoreSet(friend_root_uuid, encrypted_friendroot)
                friendroot_struct_hmac, ok15 := userlib.HMACEval(friendroot_mac_key, encrypted_friendroot)

                if ok15!=nil {return "", errors.New(strings.ToTitle("Share: cant eval friends root"))}
                userlib.DatastoreSet(friendroot_mac_uuid, friendroot_struct_hmac)
        // Form the magic string to give to the friend
            info_to_pke := append(bytes_for_uuid, friend_root_sym...)
                final_pke_info, ok16 := userlib.PKEEnc(recipients_pubkey, info_to_pke)
                 if ok16!=nil {return "", errors.New(strings.ToTitle("Share: cant pke string"))}

                signed_magic, ok17 := userlib.DSSign(userdata.SignPriv, final_pke_info)
                if ok17!=nil {return "", errors.New(strings.ToTitle("Share: cant sign a string"))}

                magic := hex.EncodeToString(final_pke_info)

                magic_uuid_bytes, ok177 := hex.DecodeString(magic)
                //magic_uuid_bytes_reduced := magic_uuid_bytes[:16]

                if ok177!=nil {return "",errors.New(strings.ToTitle("Recieve: Cant decode string"))}



                magic_uuid_bytes_extend := userlib.Argon2Key([]byte(recipient), magic_uuid_bytes, 16)

                signature_uuid, ok18 := uuid.FromBytes(magic_uuid_bytes_extend)

                userlib.DatastoreSet(signature_uuid, signed_magic)

                if ok18!=nil {return "",ok18}


            return magic, nil
}
// Note recipient's filename can be different from the sender's filename.
// The recipient should not be able to discover the sender's view on
// what the filename even is!  However, the recipient must ensure that
// it is authentically from the sender.
func (userdata *User) ReceiveFile(filename string, sender string,
        magic_string string) error {
    // Get the users current files from the datastore
                var files Files
                enc_files_bytes, ok1 := userlib.DatastoreGet(userdata.Files_uuid)

                if !ok1 {return errors.New(strings.ToTitle("Recieve: cant get files from id"))}
                dec_files_bytes := userlib.SymDec(userdata.Files_symkey, enc_files_bytes)
                json.Unmarshal(dec_files_bytes, &files)
                filename_hash, ok2 := userlib.HMACEval(files.Filename_key, []byte(filename))

                if ok2!=nil {return errors.New(strings.ToTitle("Recieve: cant hmac filename"))}
                filename_hash_string := hex.EncodeToString(filename_hash)

                if  _, ok3 := files.MyFiles[filename_hash_string]; ok3 {
                    return errors.New(strings.ToTitle("Recieve: Already have this file"))
                }
    // Decrypt with your private key and grab the info for root
            pke_msg, ok4 := hex.DecodeString(magic_string)
                if ok4!=nil {return errors.New(strings.ToTitle("Recieve: Cant decode string"))}


            pke_msg_extend := userlib.Argon2Key([]byte(userdata.Username), pke_msg, 16)

            signature_uuid, ok18 := uuid.FromBytes(pke_msg_extend[:16])

            if ok18!=nil {return errors.New(strings.ToTitle("Recieve: Cant get bytes from uuid"))}

            signature, ok33 := userlib.DatastoreGet(signature_uuid)

                if !ok33 {return errors.New(strings.ToTitle("Receive: cannot retrieve signature."))}


            verification_key, ok11 := userlib.KeystoreGet(sender + "vk")
                if !ok11 {return errors.New(strings.ToTitle("Share: without recipients pubkey"))}



            signature_verification := userlib.DSVerify(verification_key, pke_msg ,signature)

            if signature_verification != nil {

                return errors.New(strings.ToTitle("Receive: Invalid Signature."))

            }

            decrypted_magic, ok5 := userlib.PKEDec(userdata.Priv, pke_msg)

                if ok5!=nil {return errors.New(strings.ToTitle("Recieve: Cant decrypt"))}
        root_uuid := decrypted_magic[:16]
        newroot_uuid, ok6 := uuid.FromBytes(root_uuid)

                if ok6!=nil {return errors.New(strings.ToTitle("Recieve: Cant get bytes from uuid"))}
        newroot_symkey := decrypted_magic[16:32]
        files.MyFiles[filename_hash_string] = newroot_uuid
        files.FilesSymmetric[newroot_uuid] = newroot_symkey
    // Make sure the shared file/root is actually in the datastore
            _, ok7 := userlib.DatastoreGet(newroot_uuid)
        if !ok7 {return errors.New(strings.ToTitle("Recieve: Cant recieve root from datastore"))}
      //Encrypt the newest version of the users files with the corresponding hmac
                files_in_bytes, ok8 := json.Marshal(files)

                if ok8!=nil {return errors.New(strings.ToTitle("Recieve: cant marshal the files"))}
                iv1 := userlib.RandomBytes(16)
                encrypted_updated_files := userlib.SymEnc(userdata.Files_symkey, iv1, files_in_bytes)
                userlib.DatastoreSet(userdata.Files_uuid, encrypted_updated_files)
                files_updated_hmac, ok9 := userlib.HMACEval(userdata.Files_mackey, encrypted_updated_files)

                if ok9!=nil {return errors.New(strings.ToTitle("Recieve: cant hmac eval updated file"))}
                userlib.DatastoreSet(userdata.Files_macuuid, files_updated_hmac)
        return nil
}
// Removes target user's access.
func (userdata *User) RevokeFile(filename string, target_username string) (err error) {
     // Get my own files from the data store
            var files Files
            enc_files_bytes, ok1 := userlib.DatastoreGet(userdata.Files_uuid)

            if !ok1 {return errors.New(strings.ToTitle("Revoke: no file struct found"))}
            dec_files_bytes := userlib.SymDec(userdata.Files_symkey, enc_files_bytes)
            stored_files_hmac, ok2 := userlib.DatastoreGet(userdata.Files_macuuid)

            if !ok2 {return errors.New(strings.ToTitle("Revoke: cant get file bytes"))}
            compute_files_hmac, ok3 := userlib.HMACEval(userdata.Files_mackey, enc_files_bytes)

            if ok3!=nil {return errors.New(strings.ToTitle("Revoke: cant eval file bytes"))}

            if !userlib.HMACEqual(stored_files_hmac, compute_files_hmac) {
                return errors.New(strings.ToTitle("Revoke: Files structure is corrupted"))}
            json.Unmarshal(dec_files_bytes, &files)
            filename_hash, ok4 := userlib.HMACEval(files.Filename_key, []byte(filename))

            if ok4!=nil {return errors.New(strings.ToTitle("Revoke: cant eval filename"))}
            filename_hash_string := hex.EncodeToString(filename_hash)
    //Check if the file exists in the Datastore
            if  _, ok5 := files.MyFiles[filename_hash_string]; !ok5 {
                return errors.New(strings.ToTitle("Revoke: Nothing stored to revoke"))}
    // Grab the root for the file
            this_file_uuid := files.MyFiles[filename_hash_string]
            this_file_sym := files.FilesSymmetric[this_file_uuid]
            var my_root Root
            encrypted_root, ok6 := userlib.DatastoreGet(this_file_uuid)

            if !ok6 {return errors.New(strings.ToTitle("Revoke: no root in datastore"))}
            decrypted_root := userlib.SymDec(this_file_sym, encrypted_root)
            json.Unmarshal(decrypted_root, &my_root)
            stored_root_hmac, ok7 := userlib.DatastoreGet(my_root.Mac_id)

            if !ok7 {return errors.New(strings.ToTitle("Revoke: can get root mac id"))}
            compute_root_hmac, ok8 := userlib.HMACEval(my_root.RootMacKey, encrypted_root)

            if ok8!=nil {return errors.New(strings.ToTitle("Revoke: cant hmac encrypted root"))}

            if !userlib.HMACEqual(stored_root_hmac, compute_root_hmac) {
                return errors.New(strings.ToTitle("Revoke: Root structure is corrupted"))}
    // Grab the bytes of the file I want to revoke from a user
            bytes_of_file, ok9 := userdata.LoadFile(filename)

            if ok9!=nil {errors.New(strings.ToTitle("Revoke: File was manipulated loading during Revoke"))}
 // Remove the target_users info to the users file
        target_uuid := my_root.Friends_uuid[target_username]
      userlib.DatastoreDelete(target_uuid)
// Remove the target_user from the users sharedfiles
      delete(my_root.Friends_uuid, target_username)
            delete(my_root.Friends_sym, target_uuid)
// Save the new friends maps
            saved_friends_uuid := my_root.Friends_uuid
            saved_friends_sym := my_root.Friends_sym
// Delete all relevant information so I can store the file under different info
            delete(files.MyFiles, filename_hash_string)
            delete(files.FilesSymmetric, this_file_uuid)
            userlib.DatastoreDelete(this_file_uuid)
            userlib.DatastoreDelete(my_root.Start_id)
            userlib.DatastoreDelete(my_root.Mac_id)
//Update the current user Files struct in order to store new bytes
            files_in_bytes, ok10 := json.Marshal(files)

            if ok10!=nil {errors.New(strings.ToTitle("Revoke: cant marshal the files"))}
            iv1 := userlib.RandomBytes(16)
            encrypted_updated_files := userlib.SymEnc(userdata.Files_symkey, iv1, files_in_bytes)
            userlib.DatastoreSet(userdata.Files_uuid, encrypted_updated_files)
            files_updated_hmac, ok11 := userlib.HMACEval(userdata.Files_mackey, encrypted_updated_files)

            if ok11!=nil {return errors.New(strings.ToTitle("Revoke: cant eval updated file"))}
            userlib.DatastoreSet(userdata.Files_macuuid, files_updated_hmac)
// Store the file back into the user
            userdata.StoreFile(filename, bytes_of_file)
// Get the most recent file struct and root for the file
            var myfiles Files
            enc_files_bytes, ok12 := userlib.DatastoreGet(userdata.Files_uuid)

            if !ok12 {return errors.New(strings.ToTitle("Revoke: cant get the user file id"))}
            dec_files_bytes = userlib.SymDec(userdata.Files_symkey, enc_files_bytes)
            json.Unmarshal(dec_files_bytes, &myfiles)
// Check if this file even exists after it is stored
            if  _, ok13 := myfiles.MyFiles[filename_hash_string]; !ok13 {
                return errors.New(strings.ToTitle("Revoke: can find stored file"))
            }
            myfiles_uuid := myfiles.MyFiles[filename_hash_string]
            myfiles_sym := myfiles.FilesSymmetric[myfiles_uuid]
// Get the root of the file to store
            var myroot Root
            encrypted_root, ok14 := userlib.DatastoreGet(myfiles_uuid)

            if !ok14 {return errors.New(strings.ToTitle("Revoke: cant find new root"))}
            decrypted_root = userlib.SymDec(myfiles_sym, encrypted_root)
            json.Unmarshal(decrypted_root, &myroot)
            myroot.Friends_uuid = saved_friends_uuid
            myroot.Friends_sym = saved_friends_sym
// Update the roots for all my friends datastore
        for _, uuid := range myroot.Friends_uuid {
        //Get my friends root and update values
                var friend_root Root
                encrypted_root, ok15 := userlib.DatastoreGet(uuid)

                                if !ok15 {return errors.New(strings.ToTitle("Revoke: cant get friend uuid"))}
                                friend_symkey := myroot.Friends_sym[uuid]
                decrypted_root := userlib.SymDec(friend_symkey, encrypted_root)
                    json.Unmarshal(decrypted_root, &friend_root)
                friend_root.Start_id = myroot.Start_id
                                friend_root.Start_macid = myroot.Start_macid
                                friend_root.Start_symkey = myroot.Start_symkey
                                friend_root.Start_mackey = myroot.Start_mackey
                friend_root.NodeSymKey = myroot.NodeSymKey
        // encrypt the root and store it along with its hmac
                    friend_mac_uuid := friend_root.Mac_id
                friend_mackey := friend_root.RootMacKey
                iv := userlib.RandomBytes(16)
                friend_root_data, ok16 := json.Marshal(friend_root)

                                if ok16!=nil {errors.New(strings.ToTitle("Revoke: cant marshal friend root"))}
                friend_root_cipher := userlib.SymEnc(friend_symkey, iv, friend_root_data)
                userlib.DatastoreSet(uuid, friend_root_cipher)
                friend_root_HMAC, ok17 := userlib.HMACEval(friend_mackey, friend_root_cipher)

                                if ok17!=nil {return errors.New(strings.ToTitle("Revoke: cant eval friends root"))}
                userlib.DatastoreSet(friend_mac_uuid, friend_root_HMAC)
        }
//update the the users root and mac value
                myroot_mac_key := myroot.RootMacKey
                myroot_mac_uuid := myroot.Mac_id
                myroot_bytes, ok18 := json.Marshal(myroot)

                if ok18!=nil {errors.New(strings.ToTitle("Revoke: cant marshal my root"))}
                iv2 := userlib.RandomBytes(16)
                encrypted_myroot := userlib.SymEnc(myfiles_sym, iv2, myroot_bytes)
                userlib.DatastoreSet(myfiles_uuid, encrypted_myroot)
                myroot_struct_hmac, ok19 := userlib.HMACEval(myroot_mac_key, encrypted_myroot)

                if ok19!=nil {return errors.New(strings.ToTitle("Revoke: cant eval my encrypted root"))}
                userlib.DatastoreSet(myroot_mac_uuid, myroot_struct_hmac)
        return
}
