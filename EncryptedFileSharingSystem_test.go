package EncryptedFileSharingSystem

// You MUST NOT change what you import.  If you add ANY additional
// imports it will break the autograder, and we will be Very Upset.

import (
	"testing"
	"reflect"
	"github.com/cs161-staff/userlib"
	_ "encoding/json"
	_ "encoding/hex"
	_ "github.com/google/uuid"
	_ "strings"
	_ "errors"
	_ "strconv"
)

func clear() {
	// Wipes the storage so one test does not affect another
	userlib.DatastoreClear()
	userlib.KeystoreClear()
}

func TestInit1(t *testing.T) {
	clear()
	t.Log("Initialization test")

	// You can set this to false!
	userlib.SetDebugStatus(true)

	_, err := InitUser("alice", "fubar")
	if err != nil {
		// t.Error says the test fails
		t.Error("Failed to initialize user", err)
		return
	}
	// t.Log() only produces output if you run with "go test -v"
	//t.Log("Got user", u)


	//Initializing a user with no password

	u2, err := InitUser("alice", "")
	if err == nil {
		// t.Error says the test fails
		t.Error("Iniliazed a user with no password: this behavior is not allowed.", err)
		return
	}

t.Log("Got user", u2)}

	func TestInit2(t *testing.T) {
	clear()
	t.Log("Initialization test")

	// You can set this to false!
	userlib.SetDebugStatus(true)


	_, err := InitUser("", "fubar")
	if err == nil {
		// t.Error says the test fails
		t.Error("Iniliazed a user with no username: this behavior is not allowed.", err)
		return
	}

	//t.Log("Got user", u3)


}

	func TestInit3(t *testing.T) {
	clear()
	t.Log("Initialization test")

	// You can set this to false!
	userlib.SetDebugStatus(true)

	_, err := InitUser("alice", "fubar")
	if err != nil {
		// t.Error says the test fails
		t.Error("Failed to initialize user", err)
		return
	}
	// t.Log() only produces output if you run with "go test -v"
	//t.Log("Got user", u)

	//Initializing a user with the same username

	_, err1 := InitUser("alice", "fubar")
	if err1 != nil {
	t.Error("Iniliazed a user with a username that already exists: this behavior is not allowed.", err)
	      return }


	//t.Log("Got user", u1)


	// If you want to comment the line above,
	// write _ = u here to make the compiler happy
	// You probably want many more tests here.
}

func TestGetUser1(t *testing.T){
	clear()

	//Initializing two users: Alice and Bob

	u1, err := InitUser("alice", "first")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	u2, err := InitUser("bob", "second")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	v1 := []byte("This is a test")
	u1.StoreFile("file1", v1)
	v2 := []byte("This is a test")
	u2.StoreFile("file1", v2)

	// Logging in three times: twice with the correct password, and once with an incorrect one

	u3, err := GetUser("alice", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	u31, err := GetUser("alice", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	v3 := []byte("This is a test")
	u3.StoreFile("file1", v3)

	v31 := []byte("This is a test")
	u31.StoreFile("file1", v31)

	u4, err := GetUser("alice", "second")
		if err != nil {

			return
		} else

		{t.Error("Logged in using the incorrect password: this behavior is not allowed.", err)
	      return }

	 v4 := []byte("This is a test")
	u4.StoreFile("file1", v4)
	  }

	 func TestGetUser2(t *testing.T){
	clear()

	//Initializing two users: Alice and Bob

	_, err := InitUser("alice", "first")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	 // Logging in with an empty password

	 u5, err := GetUser("alice", "")
		if err != nil {
			return
		} else

		{t.Error("Logged in with no password: this behavior is not allowed.", err)
	      return }

	 // Creating another instance of the user Alice


	u6, err := GetUser("alice", "first")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	// Getting a user that has not been initialized

	u7, err := GetUser("Ryan", "first")
		if err != nil {
			return
		} else

		{t.Error("Logged in with with a user that has not been initialized: this behavior is not allowed.", err)
	      return }


	v5 := []byte("This is a test")
	u5.StoreFile("file1", v5)

	v6 := []byte("This is a test")
	u6.StoreFile("file1", v6)

	v7 := []byte("This is a test")
	u7.StoreFile("file1", v7)

}

func TestStorage0(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	u2, err := InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	v := []byte("This is a test")
	u.StoreFile("file1", v)

	v1 := []byte("This is a test: update1.")
	u.StoreFile("file1", v1)

	v11 := []byte("This is a test: update2.")
	u.StoreFile("file1", v11)

	v2 := []byte("This is a test: update2.")
	u.StoreFile("file2", v2)


	v3, err2 := u.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to upload and download", err2)
		return
	}


	if !reflect.DeepEqual(v3, v11) {
		t.Error("Downloaded file is not the same", v3, v11)
		return
	}

	err8 := u.RevokeFile("file1", "alice")
	if err8 != nil {
		t.Error("Failed to revoke alice's access to her file.", err8)
		return

	}

	magic_string1, err := u.ShareFile("file2", "bob")
	if err != nil {
		t.Error("failed to share.", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string1)
	if err != nil {
		t.Error("Failed to receive a file with from alice.", err)
		return
	}


}

func TestStorage1(t *testing.T) {
	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	_, err = InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)
	//t.Log("Got user", u2)

	// Get users Alice and Bob

	uu, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	_, err = GetUser("bob", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	// Simple storage

	v := []byte("This is a test")
	uu.StoreFile("file1", v)
	v2 := []byte("This is a testddd")
    //u.AppendFile("file1",v2)
	v2, err2 := uu.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to upload and download", err2)
		return
	}


	if !reflect.DeepEqual(v, v2) {
		t.Error("Downloaded file is not the same", v, v2)
		return
	}

	// Update the content of file 1 by calling store again with different bytes

	v21 := []byte("This is file1 after updating its content.")
	uu.StoreFile("file1", v21)

	v211, err2 := uu.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download", err2)
		return
	}

	if !reflect.DeepEqual(v211, v21) {
		t.Error("Downloaded file is not the same: failed to update the file with new content.", v211, v21)
		return
	} }

	//------------------////------------------------------------------------------
	func TestStorage2(t *testing.T) {
	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	_, err = InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)
	//t.Log("Got user", u2)

	// Get users Alice and Bob

	uu, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	uu2, err := GetUser("bob", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	// Simple storage

	v := []byte("This is a test")
	uu.StoreFile("file1", v)
	//v2 := []byte("This is a testddd")

	// Loading a file that does not exist


	_, err2 := uu2.LoadFile("file1")
	if err2 == nil {
		t.Error("Loaded a file that does not exist: this behavior should not be allowed.", err2)
		return
	}

	// Storing the file by the second user


	v3 := []byte("This is a test")
	uu2.StoreFile("file1", v3)

	v3, err3 := uu2.LoadFile("file1")
	if err3 != nil {
		t.Error("Failed to upload and download", err3)
		return
	} }

	//-------------------------------////---------------------

	func TestStorage3(t *testing.T) {
	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	_, err = InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)
	//t.Log("Got user", u2)

	// Get users Alice and Bob

	uu, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	uu2, err := GetUser("bob", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	// Share a file and then store it again with different content, then check whether it remains shared

	v := []byte("This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.")
	uu2.StoreFile("file1", v)

	magic_string1, err := uu2.ShareFile("file1", "alice")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = uu.ReceiveFile("file12", "bob", magic_string1)
	if err != nil {
		t.Error("Failed to receive a file with from Bob.", err)
		return
	}

	v4 := []byte("This is a test: updated after sharing")
	uu2.StoreFile("file1", v4)

	v2111, err2 := uu.LoadFile("file12")
	if err2 != nil {
		t.Error("Failed to download", err2)
		return
	}

	if !reflect.DeepEqual(v2111, v4) {
		t.Error("Downloaded file is not the same: failed to update the file with new content by Alice.", v2111, v4)
		return
	}

}

func TestStorage4(t *testing.T) {
	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	_, err = InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)
	//t.Log("Got user", u2)

	// Get users Alice and Bob

	uu, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	_, err = GetUser("bob", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	// Simple storage

	v := []byte("This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.This is a test.")
	uu.StoreFile("file1", v)
	v2 := []byte("This is a testddd")
    //u.AppendFile("file1",v2)
	v2, err2 := uu.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to upload and download", err2)
		return
	}


	if !reflect.DeepEqual(v, v2) {
		t.Error("Downloaded file is not the same", v, v2)
		return
	}}

//------------------------------

func TestInvalidFile1(t *testing.T) {
	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	_, err = InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u2)

	_, err = InitUser("carol", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u3)

	// Get users Alice and Bob

	uu1, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	uu2, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	uu3, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	_, err2 := uu1.LoadFile("this_file_does_not_exist")
	if err2 == nil {
		t.Error("Downloaded a ninexistent file", err2)
		return
	}


	v1 := []byte("This is a test")
	uu1.StoreFile("file1", v1)


	uu1load, err2 := uu1.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	uu2load, err2 := uu2.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	uu3load, err2 := uu3.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	if (len(v1) != len(uu1load)) {t.Error("Downloaded a different file.", err)
		return}

	if (len(v1) != len(uu2load)) {t.Error("Downloaded a different file.", err)
		return}

	if (len(v1) != len(uu3load)) {t.Error("Downloaded a different file.", err)
		return} }

	//-----------------------////--------------------------------------------------

	func TestInvalidFile2(t *testing.T) {
	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	_, err = InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u2)

	_, err = InitUser("carol", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u3)

	// Get users Alice and Bob

	uu1, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	uu2, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	uu3, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

		v1 := []byte("This is a test")
	uu1.StoreFile("file1", v1)


	v1append := []byte("first append")
	v2append := []byte("second append")
	v3append := []byte("third append")

	err = uu1.AppendFile("file1", v1append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = uu2.AppendFile("file1", v2append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = uu3.AppendFile("file1", v3append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}


	uu1load_append, err2 := uu1.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	uu2load_append, err2 := uu2.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	uu3load_append, err2 := uu3.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	if !reflect.DeepEqual(uu3load_append, uu1load_append) {t.Error("Users downloaded different files.", err)
		return}

	if !reflect.DeepEqual(uu3load_append, uu2load_append) {t.Error("Users downloaded different files.", err)
		return} }

	//------------------------////---------------------------------------

	func TestInvalidFile3(t *testing.T) {
	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)


	// Get users Alice and Bob

	uu1, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}


	v11 := []byte("")
	uu1.StoreFile("file2", v11)


	uu1_load, err2 := uu1.LoadFile("file2")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	if (len(uu1_load) != 0) {t.Error("Did not download the correct file.", err)
		return} }

	//-----------------------////------------------------------------------

	func TestInvalidFile4(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	u2, err := InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u2)

	//u3, err := InitUser("carol", "fubar")
	//if err != nil {
	//	t.Error("Failed to initialize user", err)
	//	return
	//}

	//t.Log("Got user", u3)

	// Get users Alice and Bob

	uu1, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	uu2, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}

	_, err = GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}


	v11 := []byte("Test file.")
	uu1.StoreFile("file2", v11)

	v12 := []byte("Test file.")
	uu2.StoreFile("file1", v12)


	magic_string1, err := u.ShareFile("file2", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string1)
	if err != nil {
		t.Error("failed to. receive.", err)
		return
	}

	magic_string2, err := u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file3", "alice", magic_string2)
	if err != nil {
		t.Error("Failed to receive.", err)
		return
	}

	u2_load, err2 := u2.LoadFile("file3")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	if !reflect.DeepEqual(u2_load, v11) {
		t.Error("Downloaded file is not the same.", u2_load, v11)
		return
	} }

	//------------------------------------------------------

	func TestInvalidFile5(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	u2, err := InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u2)

	v11 := []byte("Test file.")
	u.StoreFile("file1", v11)
	u.StoreFile("file3", v11)
	u2.StoreFile("file1", v11)

	magic_string1, err := u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string1)
	if err == nil {
		t.Error("received a file with the same name.", err)
		return
	}

	//-----------------------

	magic_string2, err := u.ShareFile("file3", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file3", "alice", magic_string2)
	if err != nil {
		t.Error("failed to receive a file.", err)
		return
	}



	err8 := u.RevokeFile("file3", "bob")
	if err8 != nil {
		t.Error("Failed to revoke bob's access to file1", err8)
		return

	}

	_, err2 := u2.LoadFile("file3")
	if err2 == nil {
		t.Error("Download the file after revoke.", err2)
		return
	} }

	//------------------------------------------------------

	func TestInvalidFile6(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	//u2, err := InitUser("bob", "fubar")
	//if err != nil {
	//	t.Error("Failed to initialize user", err)
	//	return
	//}

	//t.Log("Got user", u2)

	v11 := []byte("Test file.")
	u.StoreFile("file1", v11)

	v3append := []byte("Test file.")


	err = u.AppendFile("file1", v3append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	_, err2 := u.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file after append.", err2)
		return
	}}

	//-------------------------------------------------------

	func TestInvalidFile7(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	u2, err := InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u2)

	v11 := []byte("Test file.")
	u.StoreFile("file1", v11)

	v12 := []byte("Test file.")
	u2.StoreFile("file1", v12)

	magic_string1, err := u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string1)
	if err == nil {
		t.Error("received a file with the same name.", err)
		return
	}}

	//---------------------------------------

	func TestInvalidFile8(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	u2, err := InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u2)

	v11 := []byte("Test file.")
	u.StoreFile("file1", v11)

	v12 := []byte("Test file.")
	u2.StoreFile("file2", v12)

	magic_string1, err := u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}


	err = u2.ReceiveFile("file1", "alice", magic_string1)
	if err != nil {
		t.Error("Failed to receive.", err)
		return
	}

	err8 := u.RevokeFile("file1", "bob")
	if err8 != nil {
		t.Error("Failed to revoke bob's access to file1", err8)
		return

	}

	//v11 = []byte("Updated test file.")
	//u2.StoreFile("file3", v11)


	u2_load_revoke, err2 := u2.LoadFile("file1")
	if err2 == nil {
		t.Error("Successfully load after revoke.", err2)
		return
	}

	if reflect.DeepEqual(u2_load_revoke, v11) {
		t.Error("Downloaded file is the same:  this behavior is not allowed.", u2_load_revoke, v11)
		return
	}}

	//-----------------------------------------

	func TestInvalidFile9(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	u2, err := InitUser("bob", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u2)

	u3, err := InitUser("carol", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u3)

	v11 := []byte("Test file.")
	u.StoreFile("file1", v11)

	v12 := []byte("Test file.")
	u2.StoreFile("file33", v12)


	magic_string12, err := u2.ShareFile("file33", "alice")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u.ReceiveFile("file33", "bob", magic_string12)
	if err != nil {
		t.Error("Failed to recceive a file.", err)
		return
	}

	err = u3.ReceiveFile("file33", "bob", magic_string12)
	if err == nil {
		t.Error("stolen token works.", err)
		return
	}

	v44 := []byte("carol's file.")
	u3.StoreFile("file33", v44)


	magic_string44, err := u3.ShareFile("file33", "carol")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u.ReceiveFile("file333", "bob", magic_string44)
	if err == nil {
		t.Error("Attack succeeded.", err)
		return
	}

	err = u.ReceiveFile("file33", "carol", magic_string12)
	if err == nil {
		t.Error("Attack succeeded.", err)
		return
	}



}


func TestShare1(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v := []byte("This is a test")
	u.StoreFile("file1", v)

	v1 := []byte("This is a test")
	u.StoreFile("file2", v1)



	var v2 []byte
	var magic_string string

	v, err = u.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file from alice", err)
		return
	}

	magic_string, err = u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	magic_string1, err := u.ShareFile("file2", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}


	magic_string2, err := u.ShareFile("file1", "amit")
	if err == nil {
		t.Error("Shared a file with a user that does not exist: this behavior is not allowed.", err)
		return
	}

	magic_string3, err := u.ShareFile("file9", "bob")
	if err == nil {
		t.Error("Shared a file that does not exist with Bob: this behavior is not allowed.", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}


	err = u2.ReceiveFile("file1", "alice", magic_string1)
	if err == nil {
		t.Error("Received another file with the same name: this behavior is not allowed.", err)
		return
	}


	err = u2.ReceiveFile("file2", "alice", magic_string2)
	if err == nil {
		t.Error("Received a file using an invalid token sent to amit: this behavior is not allowed.", err)
		return
	}


	v2, err = u2.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file after sharing", err)
		return
	}
	if !reflect.DeepEqual(v, v2) {
		t.Error("Shared file is not the same", v, v2)
		return
	}

	err = u2.ReceiveFile("file9", "alice", magic_string3)
	if err == nil {
		t.Error("Received a file using an invalid token: this behavior is not allowed.", err)
		return
	}

}

func TestShare2(t *testing.T) {
	clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v := []byte("")
	u.StoreFile("file1", v)

	v1 := []byte("This is a test")
	//u.StoreFile("file2", v1)



	//var v2 []byte
	var magic_string string

	v, err = u.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file from alice", err)
		return
	}

	magic_string, err = u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	v, err = u2.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the empty file from alice", err)
		return
	}

	err = u2.ReceiveFile("file2", "alice", magic_string)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	err = u2.AppendFile("file1", v1)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = u.AppendFile("file1", v1)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	v = []byte("")
	u.StoreFile("file1", v)

	v12, err10 := u2.LoadFile("file1")
	if err10 != nil {
		t.Error("Failed to load after append.", err10)
		return
	}

	if !reflect.DeepEqual(v12, v) {
		t.Error("Downloaded files are different.", v12, v)
		return
	}



	}

	func TestShare3(t *testing.T) {
	clear()
	u1, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v := []byte("")
	u1.StoreFile("file1", v)
	u1.StoreFile("file1", v)
	u2.StoreFile("file1", v)

	magic_string1, err4 := u1.ShareFile("file1", "alice")
	if err4 != nil {
		t.Error("Failed to share the a file", err4)
		return
	}

	err5 := u1.ReceiveFile("file2", "alice", magic_string1)
	if err5 != nil {
		t.Error("Failed to receive a file from Alice.", err5)
		return
	}

	err6 := u1.ReceiveFile("file1", "bob", magic_string1)
	if err6 == nil {
		t.Error("Successfully got access to Bob's file using a forged access token.", err6)
		return
	}


}








func TestDeepShare1(t *testing.T) {
	clear()
	u1, err1 := InitUser("alice", "fubar")
	if err1 != nil {
		t.Error("Failed to initialize user", err1)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}


	u3, err3 := InitUser("carol", "foobar")
	if err3 != nil {
		t.Error("Failed to initialize bob", err3)
		return
	}


	v1 := []byte("This is a test")
	u1.StoreFile("file1", v1)

	magic_string1, err4 := u1.ShareFile("file1", "bob")
	if err4 != nil {
		t.Error("Failed to share the a file", err4)
		return
	}

	err5 := u2.ReceiveFile("file1", "alice", magic_string1)
	if err5 != nil {
		t.Error("Failed to receive a file from Alice.", err5)
		return
	}

	magic_string2, err6 := u2.ShareFile("file1", "carol")
	if err6 != nil {
		t.Error("Failed to share the a file", err6)
		return
	}

	err7 := u3.ReceiveFile("file1", "bob", magic_string2)
	if err7 != nil {
		t.Error("Failed to receive a file from Bob.", err7)
		return
	}

	//--------------------------------------------------------------------------------------------------

	err := u2.AppendFile("file1", v1)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	v12, err10 := u3.LoadFile("file1")
	if err10 != nil {
		t.Error("Failed to load after append.", err10)
		return
	}

	v11, err10 := u1.LoadFile("file1")
	if err10 != nil {
		t.Error("Failed to load after append.", err10)
		return
	}



	if len(v11) != len(v12) {
		t.Error("Failed to retrieve the version updated by another user Alice", err)
		return
	}


	//--------------------------------------------

	err = u3.AppendFile("file1", v1)
	if err != nil {

		t.Error("Failed to append to file3 by carol.", err)
		return

	}

	v111, err10 := u1.LoadFile("file1")
	if err10 != nil {
		t.Error("Failed to load after append.", err10)
		return
	}

	v112, err10 := u2.LoadFile("file1")
	if err10 != nil {
		t.Error("Failed to load after append.", err10)
		return
	}

	if len(v111) != len(v112) {
		t.Error("Failed to retrieve the version updated by another user Carol.", err)
		return
	} }

	//----------------------------////----------------------------------
func TestDeepShare2(t *testing.T) {
	clear()
	u1, err1 := InitUser("alice", "fubar")
	if err1 != nil {
		t.Error("Failed to initialize user", err1)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}


	u3, err3 := InitUser("carol", "foobar")
	if err3 != nil {
		t.Error("Failed to initialize bob", err3)
		return
	}

	v1 := []byte("This is a test")
	u1.StoreFile("file1", v1)

	magic_string1, err4 := u1.ShareFile("file1", "bob")
	if err4 != nil {
		t.Error("Failed to share the a file", err4)
		return
	}

	err5 := u2.ReceiveFile("file1", "alice", magic_string1)
	if err5 != nil {
		t.Error("Failed to receive a file from Alice.", err5)
		return
	}

	err8 := u1.RevokeFile("file1", "bob")
	if err8 != nil {
		t.Error("Failed to revoke bob's access to file1", err8)
		return

	}

	v1, err := u1.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file after revocation of other users.", err)
		return
	}

	v1, err9 := u2.LoadFile("file1")
	if err9 == nil {
		t.Error("Downloded a file after Alice revoked access to it: this behavior should not be allowed.", err9)
		return
	}

	v1, err14 := u3.LoadFile("file1")
	if err14 == nil {
		t.Error("Downloded a file after Alice revoked access to it: this behavior should not be allowed.", err14)
		return
	}

	err15 := u2.ReceiveFile("file2", "alice", magic_string1)
	if err5 != nil {
		t.Error("Successful replay of magic string.", err15)
		return
	}

	_, err142 := u2.LoadFile("file1")
	if err142 == nil {
		t.Error("Downloded a file after Alice revoked access to it: this behavior should not be allowed.", err14)
		return
	}



	err16 := u2.AppendFile("file2", v1)
	if err16 == nil {

		t.Error("Append after revoke: this behavior should not be allowed.", err16)
		return

	} }

	//-----------------------------------------////---------
	func TestDeepShare3(t *testing.T) {
	clear()

	u1, err1 := InitUser("alice", "fubar")
	if err1 != nil {
		t.Error("Failed to initialize user", err1)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}


	v1 := []byte("This is a test")
	u1.StoreFile("file1", v1)



	u13, err13 := InitUser("attacker", "fubar")
	if err13 != nil {
		t.Error("Failed to initialize user", err13)
		return
	}

	v2 := []byte("This is a test")
	u1.StoreFile("file14", v2)

	v13 := []byte("This is a test")
	u13.StoreFile("file13", v13)

	magic_string14, err14 := u1.ShareFile("file14", "bob")
	if err14 != nil {
		t.Error("Failed to share the a file", err14)
		return
	}

	magic_string13, err133 := u13.ShareFile("file13", "bob")
	if err133 != nil {
		t.Error("Failed to share the a file", err133)
		return
	}

	err1555 := u13.ReceiveFile("file14", "alice", magic_string14)
	if err1555 == nil {
		t.Error("Attacker received bob's file: this behavior should not be allowed.", err1555)
		return
	}

	err155 := u2.ReceiveFile("file14", "alice", magic_string13)
	if err155 == nil {
		t.Error("Received file from the attacker instead: this behavior should not be allowed.", err155)
		return
	}


	}


func TestAppend(t *testing.T) {

	clear()
	_, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}

	//t.Log("Got user", u)

	// Get users Alice and Bob

	uu, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get another instance of the user", err)
			return
		}


	v := []byte("This is a test")
	uu.StoreFile("file1", v)


	v, err = uu.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file from alice", err)
		return
	}

	v1 := []byte("Content to be appended to file1")



	err = uu.AppendFile("file1", v1)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = uu.AppendFile("file3", v1)
	if err == nil {

		t.Error("Successfully appended to a file that does not exist: this behavior should not be allowed.", err)
		return

	}

}



func TestRevoke1(t *testing.T) {
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v := []byte("This is a test")
	u.StoreFile("file1", v)

	v1 := []byte("This is a test")
	u.StoreFile("file2", v1)

	//var v2 []byte
	var magic_string string

	v, err = u.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file from alice", err)
		return
	}

	magic_string, err = u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}
	err = u2.ReceiveFile("file2", "alice", magic_string)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	err = u.RevokeFile("file1", "bob")
	if err != nil {
		t.Error("Failed to revoke bob's access to file1", err)
		return

	}

	v, err = u2.LoadFile("file1")
	if err == nil {
		t.Error("Bob loaded the file that he is not supposed to access after Alice revoked his access to it", err)
		return
	} }

func TestRevoke2(t *testing.T) {
clear()
u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v := []byte("This is a test")
	u.StoreFile("file1", v)

	v1 := []byte("This is a test")
	u.StoreFile("file2", v1)

	var magic_string string

	magic_string, err = u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}



	err = u2.ReceiveFile("file2", "alice", magic_string)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	err = u.RevokeFile("file1", "bob")
	if err != nil {
		t.Error("Failed to revoke bob's access to file1", err)
		return

	}

		v2 := []byte("Content to be appended to file1")


	err = u2.AppendFile("file1", v2)
	if err == nil {

		t.Error("Successfully appended to a file for which the user is no longer allowed to access: this behavior is not allowed.", err)
		return

	}


	err = u.RevokeFile("file3", "bob")
	if err == nil {
		t.Error("Revoked access to a file that does not exist (not shared) with Bob: this behavior should not be allowed.", err)
		return

	}


}

func TestRevoke3(t *testing.T) {
clear()
	u1, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v := []byte("This is a test")
	u1.StoreFile("file1", v)

	v1 := []byte("This is a test")
	u2.StoreFile("file2", v1)

	err = u1.RevokeFile("file2", "bob")
	if err == nil {
		t.Error("Successfully revoked Bob's access to his file.", err)
		return

	}
}


func TestActiveUsers1(t *testing.T){
clear()

	u1, err := InitUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	//u2, err := InitUser("alice", "second")
		//if err != nil {
		//	t.Error("Failed to initialize user", err)
			//return
		//}

	u11, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	u12, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	v := []byte("This is a test")

	u1.StoreFile("file1", v)

	v, err = u11.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Alice", err)
		return
	}

	v, err = u12.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file from alice by another active instantiation of Alice (u2)", err)
		return
	} }

	func TestActiveUsers2(t *testing.T){
clear()

	u1, err := InitUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	//u2, err := InitUser("alice", "second")
		//if err != nil {
		//	t.Error("Failed to initialize user", err)
			//return
		//}

	u11, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	u12, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	v := []byte("This is a test")

	u1.StoreFile("file1", v)


	v1 := []byte("Content to be appended to file1")

	err = u11.AppendFile("file1", v1)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	v3, err := u11.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Alice", err)
		return
	}
	v4, err := u12.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file from alice by another active instantiation of Alice (u2)", err)
		return
	}

	if len(v3) != len(v4) {
		t.Error("Failed to retrieve the version updated by another user Alice", err)
		return
	} }

	func TestActiveUsers3(t *testing.T){
clear()

	u1, err := InitUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	//u2, err := InitUser("alice", "second")
		//if err != nil {
		//	t.Error("Failed to initialize user", err)
			//return
		//}

	u11, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	u12, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	v := []byte("This is a test")

	u1.StoreFile("file1", v)

	u3, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v5 := []byte("This is a the file to be shared")
	//v6 := []byte("This is a the file to be loaded")
	u3.StoreFile("file2", v5)

	var magic_string string

	v5, err = u3.LoadFile("file2")
	if err != nil {
		t.Error("Failed to download the file from alice", err)
		return
	}

	magic_string, err = u3.ShareFile("file2", "alice")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u11.ReceiveFile("file2", "bob", magic_string)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}


	v5, err = u12.LoadFile("file2")
	if err != nil {
		t.Error("Failed to download the file from alice that was received by another instance of user Alice", err)
		return
	} }

	//------------------------Mar18 updates---------------------------
func TestActiveUsers4(t *testing.T){
clear()

	u1, err := InitUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	//u2, err := InitUser("alice", "second")
		//if err != nil {
		//	t.Error("Failed to initialize user", err)
			//return
		//}

	//u11, err := GetUser("alice", "fubar")
	//	if err != nil {
	//		t.Error("Failed to get user", err)
	//		return
	//	}

	//u12, err := GetUser("alice", "fubar")
	//	if err != nil {
	//		t.Error("Failed to get user", err)
	//		return
	//	}

	v := []byte("This is a test")

	u1.StoreFile("file1", v)

	u4, err := InitUser("Noura", "first")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	u41, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}
	u42, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	v4test := []byte("This is a test")

	u4.StoreFile("file1", v4test)

	v41, err := u41.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	v42, err := u42.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	if (len(v41) != len(v42)) {

		t.Error("Two instances downloaded different files.", err)
		return
	} }

	//---------------------------------------

	func TestActiveUsers5(t *testing.T){
clear()

	_, err := InitUser("Noura", "first")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

		//t.Log("Got user", u4)

	u41, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}
	u42, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

		v4test := []byte("This is a test")

	u41.StoreFile("file1", v4test)

		v1 := []byte("Content to be appended to file1")

	err = u41.AppendFile("file1", v1)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	v411, err := u41.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	v421, err := u42.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	if (len(v411) != len(v421)) {

		t.Error("Two instances downloaded different files.", err)
		return
	}

	err = u42.AppendFile("file1", v4test)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	v4111, err := u41.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	v4211, err := u42.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	if (len(v4111) != len(v4211)) {

		t.Error("Two instances downloaded different files.", err)
		return
	}

	//-----------------------------

	v41test := []byte("This is a second test")

	u42.StoreFile("file2", v41test)

	v41111, err := u41.LoadFile("file2")
	if err != nil {
		t.Error("Failed to download the file uploaded by the second instance.", err)
		return
	}

	v42111, err := u42.LoadFile("file2")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	if (len(v41111) != len(v42111)) {

		t.Error("Two instances downloaded different files.", err)
		return
	} }

	//---------------------------------//-------------------------------------------//

	func TestActiveUsers6(t *testing.T){
clear()

	u1, err := InitUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	v := []byte("This is a test")

	u1.StoreFile("file2", v)

u4, err := InitUser("Noura", "first")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	u41, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}
	u42, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	magic_string41, err := u1.ShareFile("file2", "Noura")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u4.ReceiveFile("file3", "alice", magic_string41)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	v421111, err := u42.LoadFile("file3")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	v411111, err := u41.LoadFile("file3")
	if err != nil {
		t.Error("Failed to download the file received by the second instance.", err)
		return
	}

	if (len(v411111) != len(v421111)) {

		t.Error("Two instances downloaded different files.", err)
		return
	} }

	//---------------------------------//------//-------------------------------------//

func TestActiveUsers7(t *testing.T){
clear()

	u1, err := InitUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	v := []byte("This is a test")

	u1.StoreFile("file2", v)

	u4, err := InitUser("Noura", "first")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	u41, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}
	u42, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	magic_string41, err := u1.ShareFile("file2", "Noura")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u4.ReceiveFile("file2", "alice", magic_string41)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}



	err = u1.RevokeFile("file2", "Noura")
	if err != nil {
		t.Error("Failed to revoke bob's access to file1", err)
		return

	}

	v4test := []byte("This is a test")

	u1.StoreFile("file3", v4test)

	err = u1.AppendFile("file2", v4test)
	if err != nil {

		t.Error("Failed to append to file2", err)
		return

	}


	magic_string411, err := u1.ShareFile("file3", "Noura")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u41.ReceiveFile("file31", "alice", magic_string411)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	v41test := []byte("This is a test")



	err = u41.AppendFile("file31", v41test)
	if err != nil {

		t.Error("Failed to append to file3", err)
		return

	}

	v1_revoke_append, err := u42.LoadFile("file31")
	if err != nil {
		t.Error("Failed to download the file updated by the second instance.", err)
		return
	}

	v4_revoke_append, err := u41.LoadFile("file31")
	if err != nil {
		t.Error("Failed to download the file uploaded by the second instance.", err)
		return
	}

	if (len(v1_revoke_append) != len(v4_revoke_append)) {

		t.Error("Two instances did not download the same files: this behavior is incorrect.", err)
		return
	} }

	//-----more append tests -----------------------------------------

	func TestActiveUsers8(t *testing.T){
clear()

	u4, err := InitUser("Noura", "first")
		if err != nil {
			t.Error("Failed to initialize user", err)
			return
		}

	u41, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}
	u42, err := GetUser("Noura", "first")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

		v := []byte("This is a test")

	u4.StoreFile("file1", v)



	v41_append := []byte("Content to be appended to file1")
	v42_append := []byte("Content to be appended to file1")



	err = u41.AppendFile("file1", v41_append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = u42.AppendFile("file1", v42_append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	u42_append, err := u42.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Noura", err)
		return
	}

	u41_append, err := u41.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file uploaded by the second instance.", err)
		return
	}

	if (len(u42_append) != len(u41_append)) {

		t.Error("Two instances downloaded different files.", err)
		return
	}


}

func TestReceive(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize user", err)
		return
	}
	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v := []byte("This is a test")
	u.StoreFile("file1", v)

	v1 := []byte("This is a test")
	v3 := []byte("This is a test")
	u2.StoreFile("file1", v1)
	u.StoreFile("file3", v3)

	var magic_string string

	v, err = u.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file from alice", err)
		return
	}

	magic_string, err = u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	magic_string1, err1 := u.ShareFile("file3", "bob")
	if err1 != nil {
		t.Error("Failed to share the a file", err1)
		return
	}


	err = u2.ReceiveFile("file2", "alice", magic_string)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}


	err2 = u2.ReceiveFile("file1", "alice", magic_string1)
	if err2 == nil {
		t.Error("Received a file with a name that already exists in Bob's directory: this behavior is not allowed.", err2)
		return
	}


	err = u.ReceiveFile("file1", "bob", magic_string)
	if err == nil {
		t.Error("Received a file that was not shared by bob.", err)
		return
	}
}

func TestOther01(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	//u2, err2 := InitUser("bob", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize bob", err2)
	//	return
	//}

	//u3, err2 := InitUser("carol", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize carol", err2)
	//	return
	//}

	//u4, err2 := InitUser("jason", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize jason", err2)
	//	return
	//}


    //Storing 2 files in a specific order and then loading them in the reverse order

	v1 := []byte("This is a test")
	u.StoreFile("file1", v1)


	v2 := []byte("This is a test")
	u.StoreFile("file2", v2)

	v3 := []byte("This is a test")
	u.StoreFile("file3", v3)

	v2, err2 := u.LoadFile("file2")
	if err2 != nil {
		t.Error("Failed to upload and download", err2)
		return
	}

	v1, err2 = u.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to upload and download", err2)
		return
	} }

	func TestOther02(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	//u3, err2 := InitUser("carol", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize carol", err2)
	//	return
	//}

	u4, err2 := InitUser("jason", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize jason", err2)
		return
	}

	v1 := []byte("This is a test")
	u.StoreFile("file1", v1)

	v2 := []byte("This is a test")
	u.StoreFile("file2", v2)

	//Sharing 2 files and then receiving them in the reverse order

	var magic_string1 string
	var magic_string2 string
	//var magic_string3 string



	magic_string1, err = u.ShareFile("file2", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	magic_string2, err = u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}


	err = u2.ReceiveFile("file1", "alice", magic_string2)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	err = u2.ReceiveFile("file2", "alice", magic_string1)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	magic_string11, err := u2.ShareFile("file2", "jason")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u4.ReceiveFile("file2", "bob", magic_string11)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}


	//v2, err := u2.LoadFile("file1")
	//if err != nil {
	//	t.Error("Failed to download the file after sharing", err)}

	//v2, err := u2.LoadFile("file2")
	//if err != nil {
	//	t.Error("Failed to download the file after sharing", err)}
}

		func TestOther03(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	u3, err2 := InitUser("carol", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize carol", err2)
		return
	}

	v1 := []byte("This is a test")
	u.StoreFile("file1", v1)

	v2 := []byte("This is a test")
	u.StoreFile("file2", v2)

	magic_string1, err := u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string1)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	//u4, err2 := InitUser("jason", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize jason", err2)
	//	return
	//}

	// Revoke, share with a third user and then test whether the revoked user can still access the file

	err = u.RevokeFile("file1", "bob")
	if err != nil {
		t.Error("Failed to revoke bob's access to file1", err)
		return

	}

	magic_string3, err := u.ShareFile("file1", "carol")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u3.ReceiveFile("file1", "alice", magic_string3)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	err = u3.ReceiveFile("file1", "bob", magic_string3)
	if err == nil {
		t.Error("Received the file from Bob, when it should only be received from the original sender Alice: This behavior is not  allowed.", err)
		return
	}

	magic_string2, err := u.ShareFile("file2", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}


	err = u2.ReceiveFile("file2", "alice", magic_string2)
	if err != nil {
		t.Error("Failed to receive the share message", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string2)
	if err == nil {
		t.Error("Received after revoked: this behavior is not allowed.", err)
		return
	}

	//v2, err := u2.LoadFile("file1")
	//if err == nil {
		//t.Error("Downloaded a file that has been revoked: this behavior is not allowed.", err)}
}



	func TestOther04(t *testing.T){
clear()
	//u, err := InitUser("alice", "fubar")
	//if err != nil {
	//	t.Error("Failed to initialize alice", err)
	//	return
	//}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	//u3, err2 := InitUser("carol", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize carol", err2)
	//	return
	//}

	//u4, err2 := InitUser("jason", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize jason", err2)
	//	return
	//}

	// Get an access token for your files and then try to use it to access other's files

	v2_bob := []byte("This is a test file for Bob")
	u2.StoreFile("file2", v2_bob)


	v2_bob, err2 = u2.LoadFile("file2")
	if err2 != nil {
		t.Error("Failed to upload and download", err2)
		return
	}


	magic_string4, err := u2.ShareFile("file2", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file3", "alice", magic_string4)
	if err == nil {
		t.Error("Received a file when Bob is not allowed to access it by using the wrong magic string: this behavior is not allowed.", err)
		return
	} }

	func TestOther05(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	v1 := []byte("This is file1 --")
	u.StoreFile("file1", v1)

	magic_string, err := u.ShareFile("file1", "bob")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u2.ReceiveFile("file1", "alice", magic_string)
	if err != nil {
		t.Error("Failed to receive file1.", err)
		return
	}

	//u3, err2 := InitUser("carol", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize carol", err2)
	//	return
	//}

	//u4, err2 := InitUser("jason", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize jason", err2)
	//	return
	//}

	// Store a file that already exists and then test whether those you shared the file with can still access the file

	v1 = []byte("This is a test of overwriting file1")
	u.StoreFile("file1", v1)

	v2, err := u2.LoadFile("file1")
	if err != nil {
		t.Error("Failed to download the file by Bob after overwriting it by Alice", err)}


	if !reflect.DeepEqual(v2, v1) {t.Error("User downloaded a different file.", err)
		return}




	err = u.RevokeFile("file1", "bob")
	if err != nil {
		t.Error("Failed to revoke bob's access to file1", err)
		return}

	v2, err = u2.LoadFile("file2")
	if err == nil {
		t.Error("Bob downloaded file 2 when in fact he should not be allowed to do so (Bob's access to the file was revoked by Alice.", err)}

	}

  func TestOther11(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	u3, err2 := InitUser("carol", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize carol", err2)
		return
	}

	u4, err2 := InitUser("jason", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize jason", err2)
		return
	}

	v1 := []byte("This is a test")
	u.StoreFile("file1", v1)


	v2 := []byte("This is a test")
	u.StoreFile("file2", v2)

	v3 := []byte("This is a test")
	u.StoreFile("file3", v3)

	v12 := []byte("This is a test")
	u2.StoreFile("file1", v12)


	v22 := []byte("This is a test")
	u2.StoreFile("file2", v22)

	v32:= []byte("This is a test")
	u2.StoreFile("file3", v32)

	v13 := []byte("This is a test")
	u3.StoreFile("file1", v13)


	v23 := []byte("This is a test")
	u3.StoreFile("file2", v23)

	v33:= []byte("This is a test")
	u3.StoreFile("file3", v33)



	magic_string1, err := u.ShareFile("file1", "jason")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	magic_string2, err := u2.ShareFile("file1", "jason")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u4.ReceiveFile("file11", "alice", magic_string1)
	if err != nil {
		t.Error("Failed to receive a file from alice.", err)
		return
	}

	err = u4.ReceiveFile("file12", "bob", magic_string2)
	if err != nil {
		t.Error("Failed to receive a file from bob.", err)
		return
	}

	err = u.RevokeFile("file1", "jason")
	if err != nil {
		t.Error("Failed to revoke jason's access to file1", err)
		return

	}

	err = u2.RevokeFile("file1", "jason")
	if err != nil {
		t.Error("Failed to revoke jason's access to file1", err)
		return

	}

	magic_string21, err := u.ShareFile("file1", "jason")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	magic_string22, err := u.ShareFile("file2", "jason")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u4.ReceiveFile("file21", "alice", magic_string21)
	if err != nil {
		t.Error("Failed to receive a file from alice.", err)
		return
	}

	//err = u4.ReceiveFile("file22", "alice", magic_string21)
	//if err == nil {
	//	t.Error("Received incorrect file: this behavior should not be allowed.", err)
		//return
	//}

	err = u4.ReceiveFile("file22", "alice", magic_string22)
	if err != nil {
		t.Error("Failed to receive a file from alice.", err)
		return
	}

	err = u.RevokeFile("file1", "jason")
	if err != nil {
		t.Error("Failed to revoke jason's access to file1", err)
		return

	}

	err = u4.ReceiveFile("file21", "alice", magic_string21)
	if err == nil {
		t.Error("Attack succeeded.", err)
		return
	}

}

	//------------------------------------------------------------

func TestOther12(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	u3, err2 := InitUser("carol", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize carol", err2)
		return
	}

	//u4, err2 := InitUser("jason", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize jason", err2)
	//	return
	//}

	v1 := []byte("This is a test")
	u.StoreFile("file1", v1)


	v2 := []byte("This is a test")
	u.StoreFile("file2", v2)

	v3 := []byte("This is a test")
	u.StoreFile("file3", v3)

	v12 := []byte("This is a test")
	u2.StoreFile("file1", v12)


	v22 := []byte("This is a test")
	u2.StoreFile("file2", v22)

	v32:= []byte("This is a test")
	u2.StoreFile("file3", v32)

	v13 := []byte("This is a test")
	u3.StoreFile("file1", v13)


	v23 := []byte("This is a test")
	u3.StoreFile("file2", v23)

	v33:= []byte("This is a test")
	u3.StoreFile("file3", v33)

	u31_load, err2 := u3.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	u32_load, err2 := u3.LoadFile("file3")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	if !reflect.DeepEqual(u31_load, v13) {t.Error("Users downloaded different files.", err)
		return}

	if !reflect.DeepEqual(u32_load, v23) {t.Error("Users downloaded different files.", err)
		return} }


	//----------------------------------------------------------------

func TestOther13(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	u3, err2 := InitUser("carol", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize carol", err2)
		return
	}

	u4, err2 := InitUser("jason", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize jason", err2)
		return
	}

	v1 := []byte("This is a test")
	u.StoreFile("file1", v1)


	v2 := []byte("This is a test")
	u.StoreFile("file2", v2)

	v3 := []byte("This is a test")
	u.StoreFile("file3", v3)

	v12 := []byte("This is a test")
	u2.StoreFile("file1", v12)


	v22 := []byte("This is a test")
	u2.StoreFile("file2", v22)

	v32:= []byte("This is a test")
	u2.StoreFile("file3", v32)

	v13 := []byte("This is a test")
	u3.StoreFile("file1", v13)


	v23 := []byte("This is a test")
	u3.StoreFile("file2", v23)

	v33:= []byte("This is a test")
	u3.StoreFile("file3", v33)

	magic_string31, err := u3.ShareFile("file1", "alice")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}


	magic_string21, err := u.ShareFile("file1", "jason")
	if err != nil {
		t.Error("Failed to share the a file", err)
		return
	}

	err = u4.ReceiveFile("file21", "alice", magic_string21)
	if err != nil {
		t.Error("Failed to receive a file from alice.", err)
		return
	}


	err = u.ReceiveFile("file31", "carol", magic_string21)
	if err == nil {
		t.Error("Attack succeeded.", err)
		return
	}

	err = u.ReceiveFile("file31", "carol", magic_string31)
	if err != nil {
		t.Error("Failed to receive the file.", err)
		return
	} }

	//------------------------------------------------------------------

	func TestOther14(t *testing.T){
clear()
	u, err := InitUser("alice", "fubar")
	if err != nil {
		t.Error("Failed to initialize alice", err)
		return
	}


	u2, err2 := InitUser("bob", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize bob", err2)
		return
	}

	u3, err2 := InitUser("carol", "foobar")
	if err2 != nil {
		t.Error("Failed to initialize carol", err2)
		return
	}

	//u4, err2 := InitUser("jason", "foobar")
	//if err2 != nil {
	//	t.Error("Failed to initialize jason", err2)
	//	return
	//}

	v1 := []byte("This is a test")
	u.StoreFile("file1", v1)


	v2 := []byte("This is a test")
	u.StoreFile("file2", v2)

	v3 := []byte("This is a test")
	u.StoreFile("file3", v3)

	v12 := []byte("This is a test")
	u2.StoreFile("file1", v12)


	v22 := []byte("This is a test")
	u2.StoreFile("file2", v22)

	v32:= []byte("This is a test")
	u2.StoreFile("file3", v32)

	v13 := []byte("This is a test")
	u3.StoreFile("file1", v13)


	v23 := []byte("This is a test")
	u3.StoreFile("file2", v23)

	v33:= []byte("This is a test")
	u3.StoreFile("file3", v33)

	v1_append := []byte("")
	v2_append := []byte("1")
	v3_append := []byte("3")
	v4_append := []byte("4")
	v_5append := []byte("134")

	err = u.AppendFile("file1", v1_append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = u.AppendFile("file1", v2_append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = u.AppendFile("file1", v3_append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	err = u.AppendFile("file1", v4_append)
	if err != nil {

		t.Error("Failed to append to file1", err)
		return

	}

	u_load, err2 := u.LoadFile("file1")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}

	append_result := append(v_5append, v13...)

	if (len(u_load) != len(append_result)) {t.Error("Append is not working as intended.", err)}

	u_2, err := GetUser("alice", "fubar")
		if err != nil {
			t.Error("Failed to get user", err)
			return
		}

	u_2load := []byte("This is a test")

	u_2.StoreFile("fileu_2", u_2load)

	u_load, err2 = u.LoadFile("fileu_2")
	if err2 != nil {
		t.Error("Failed to download the file.", err2)
		return
	}



	if !reflect.DeepEqual(u_2load, u_load) {t.Error("Two instances downloaded different files.", err)
		return}

	err = u.RevokeFile("file1", "jason")
	if err != nil {
		t.Error("Failed to revoke jason's access to file1", err)
		return

	}

}

func TestAlotOfThings1(t *testing.T) {
        clear()
        _, err1 := InitUser("alice", "first")
        if err1 != nil {
                t.Error("Failed to initialize user", err1)
                return
        }
        _, err2 := InitUser("bob", "second")
        if err2 != nil {
                t.Error("Failed to initialize user", err2)
                return
        }
        _, err3 := InitUser("carol", "third")
        if err3 != nil {
                t.Error("Failed to initialize user", err3)
                return
        }
        alice1, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice2, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice3, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob1, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob2, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob3, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol1, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol2, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol3, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice1file1 := []byte("this is just a file to store for alice")
        alice2file2 := []byte("this will be another file stored by alice")
        alice3file3 := []byte("yet another file that alice will store")
        bob1file1 := []byte("this is just a file to store for bob")
        bob2file2 := []byte("this will be another file stored for bob")
        bob3file3 := []byte("yet another file that bob will store")
        carol1file1 := []byte("this is just a file to store for carol")
        carol2file2 := []byte("this will be another file stored for carol")
        carol3file3 := []byte("yet another file that bob will store")
        alice1.StoreFile("file1", alice1file1)
        alice2.StoreFile("file2", alice2file2)
        alice3.StoreFile("file3", alice3file3)
        bob1.StoreFile("file1", bob1file1)
        bob2.StoreFile("file2", bob2file2)
        bob3.StoreFile("file3", bob3file3)
        carol1.StoreFile("file1", carol1file1)
        carol2.StoreFile("file2", carol2file2)
        carol3.StoreFile("file3", carol3file3)
// check that alice1 can load everything from other alice2/3
        alice1loaded1, err := alice1.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice1loaded1, alice1file1) {
                t.Error("Downloaded file is not the same", alice1loaded1, alice1file1)
                return
        }
        alice1loaded2, err := alice1.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice1loaded2, alice2file2) {
                t.Error("Downloaded file is not the same", alice1loaded2, alice2file2)
                return
        }
        alice1loaded3, err := alice1.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice1loaded3, alice3file3) {
                t.Error("Downloaded file is not the same", alice1loaded3, alice3file3)
                return
        }
        // Check that alice 2 can load everything from alice1/3
        alice2loaded1, err := alice2.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice2loaded1, alice1file1) {
                t.Error("Downloaded file is not the same", alice2loaded1, alice1file1)
                return
        }
        alice2loaded2, err := alice2.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice2loaded2, alice2file2) {
                t.Error("Downloaded file is not the same", alice2loaded2, alice2file2)
                return
        }
        alice2loaded3, err := alice2.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice2loaded3, alice3file3) {
                t.Error("Downloaded file is not the same", alice2loaded3, alice3file3)
                return
        }
// Check that alice 3 can load everything from alice1/2
        alice3loaded1, err := alice3.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice3loaded1, alice1file1) {
                t.Error("Downloaded file is not the same", alice3loaded1, alice1file1)
                return
        }
        alice3loaded2, err := alice3.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice3loaded2, alice2file2) {
                t.Error("Downloaded file is not the same", alice3loaded2, alice2file2)
                return
        }
        alice3loaded3, err := alice3.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(alice3loaded3, alice3file3) {
                t.Error("Downloaded file is not the same", alice3loaded3, alice3file3)
                return
        }
  // check that bob1 can load everything from bob2/3
        bob1loaded1, err := bob1.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob1loaded1, bob1file1) {
                t.Error("Downloaded file is not the same", bob1loaded1, bob1file1)
                return
        }
        bob1loaded2, err := bob1.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob1loaded2, bob2file2) {
                t.Error("Downloaded file is not the same", bob1loaded2, bob2file2)
                return
        }
        bob1loaded3, err := bob1.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob1loaded3, bob3file3) {
                t.Error("Downloaded file is not the same", bob1loaded3, bob3file3)
                return
        }
        // Check that bob 2 can load everything from bob1/3
        bob2loaded1, err := bob2.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob2loaded1, bob1file1) {
                t.Error("Downloaded file is not the same", bob2loaded1, bob1file1)
                return
        }
        bob2loaded2, err := bob2.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob2loaded2, bob2file2) {
                t.Error("Downloaded file is not the same", bob2loaded2, bob2file2)
                return
        }
        bob2loaded3, err := bob2.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob2loaded3, bob3file3) {
                t.Error("Downloaded file is not the same", bob2loaded3, bob3file3)
                return
        }
        // Check that bob 3 can load everything from bob1/2
        bob3loaded1, err := bob3.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob3loaded1, bob1file1) {
                t.Error("Downloaded file is not the same", bob3loaded1, bob1file1)
                return
        }
        bob3loaded2, err := bob3.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob3loaded2, bob2file2) {
                t.Error("Downloaded file is not the same", bob3loaded2, bob2file2)
                return
        }
        bob3loaded3, err := bob3.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(bob3loaded3, bob3file3) {
                t.Error("Downloaded file is not the same", bob3loaded3, bob3file3)
                return
        }
        // check that carol1 can load everything from carol2/3
        carol1loaded1, err := carol1.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol1loaded1, carol1file1) {
                t.Error("Downloaded file is not the same", carol1loaded1, carol1file1)
                return
        }
        carol1loaded2, err := carol1.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol1loaded2, carol2file2) {
                t.Error("Downloaded file is not the same", carol1loaded2, carol2file2)
                return
        }
        carol1loaded3, err := carol1.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol1loaded3, carol3file3) {
                t.Error("Downloaded file is not the same", carol1loaded3, carol3file3)
                return
        }
        // Check that carol 2 can load everything from carol1/3
        carol2loaded1, err := carol2.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol2loaded1, carol1file1) {
                t.Error("Downloaded file is not the same", carol2loaded1, carol1file1)
                return
        }
        carol2loaded2, err := carol2.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol2loaded2, carol2file2) {
                t.Error("Downloaded file is not the same", carol2loaded2, carol2file2)
                return
        }
        carol2loaded3, err := carol2.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol2loaded3, carol3file3) {
                t.Error("Downloaded file is not the same", carol2loaded3, carol3file3)
                return
        }
        // Check that carol 3 can load everything from carol1/2
        carol3loaded1, err := carol3.LoadFile("file1")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol3loaded1, carol1file1) {
                t.Error("Downloaded file is not the same", carol3loaded1, carol1file1)
                return
        }
        carol3loaded2, err := carol3.LoadFile("file2")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol3loaded2, carol2file2) {
                t.Error("Downloaded file is not the same", carol3loaded2, carol2file2)
                return
        }
        carol3loaded3, err := carol3.LoadFile("file3")
        if err != nil {
                t.Error("Failed to upload and download", err)
                return
        }
        if !reflect.DeepEqual(carol3loaded3, carol3file3) {
                t.Error("Downloaded file is not the same", carol3loaded3, carol3file3)
                return
        } }

        //---------------------------------------------------------------------------------------



        func TestAlotOfThings2(t *testing.T) {
        clear()
        _, err1 := InitUser("alice", "first")
        if err1 != nil {
                t.Error("Failed to initialize user", err1)
                return
        }
        _, err2 := InitUser("bob", "second")
        if err2 != nil {
                t.Error("Failed to initialize user", err2)
                return
        }
        _, err3 := InitUser("carol", "third")
        if err3 != nil {
                t.Error("Failed to initialize user", err3)
                return
        }
        alice1, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice2, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice3, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob1, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob2, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob3, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol1, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol2, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol3, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice1file1 := []byte("this is just a file to store for alice")
        alice2file2 := []byte("this will be another file stored by alice")
        alice3file3 := []byte("yet another file that alice will store")
        bob1file1 := []byte("this is just a file to store for bob")
        bob2file2 := []byte("this will be another file stored for bob")
        bob3file3 := []byte("yet another file that bob will store")
        carol1file1 := []byte("this is just a file to store for carol")
        carol2file2 := []byte("this will be another file stored for carol")
        carol3file3 := []byte("yet another file that bob will store")
        alice1.StoreFile("file1", alice1file1)
        alice2.StoreFile("file2", alice2file2)
        alice3.StoreFile("file3", alice3file3)
        bob1.StoreFile("file1", bob1file1)
        bob2.StoreFile("file2", bob2file2)
        bob3.StoreFile("file3", bob3file3)
        carol1.StoreFile("file1", carol1file1)
        carol2.StoreFile("file2", carol2file2)
        carol3.StoreFile("file3", carol3file3)

        firstbytesforfile := []byte("The start of a shared file/")
        checkthefile := firstbytesforfile
        alice1.StoreFile("will it break", firstbytesforfile)
        secondbytesforfile := []byte("Second Instance of alice has appened to this file/")
        checkthefile = append(checkthefile, secondbytesforfile...)
        err = alice2.AppendFile("will it break", secondbytesforfile)
        if err != nil {
                t.Error("Failed to append to will it break", err)
                return
        }
        thirdbytesforfile := []byte("Third Instance of alice has appened even more/")
        checkthefile = append(checkthefile, thirdbytesforfile...)
        err = alice3.AppendFile("will it break", thirdbytesforfile)
        if err != nil {
                t.Error("Failed to append to will it break", err)
                return
        }
        alice1appendedfromload, err := alice1.LoadFile("will it break")
        if err != nil {
                t.Error("Failed to load all the appeneds", err)
        }
        if !reflect.DeepEqual(alice1appendedfromload, checkthefile) {
                t.Error("Incorrect appending", alice1appendedfromload, checkthefile)
        }

        magicstringforbob, err := alice1.ShareFile("will it break", "bob")
        if err != nil {
                t.Error("Failed to share the file with bob", err)
                return
        }
        err = bob2.ReceiveFile("break it bob", "alice", magicstringforbob)
        if err != nil {
                t.Error("received file broke for bob", err)
                return
        }
        bobfilebytes, err := bob3.LoadFile("break it bob")
        if err!=nil {
                t.Error("bob couldnt load the shared file", err)
        }
        if !reflect.DeepEqual(bobfilebytes, checkthefile) {
                t.Error("incorrect load for bob", bobfilebytes, checkthefile)
        }

        magicstringforcarol, err := alice2.ShareFile("will it break", "carol")
        if err != nil {
                t.Error("Failed to share the file with carol", err)
                return
        }
        err = carol3.ReceiveFile("break it carol", "alice", magicstringforcarol)
        if err != nil {
                t.Error("received file broke for carol", err)
                return
        }
        carolfilebytes, err := carol1.LoadFile("break it carol")
        if err!=nil {
                t.Error("carol couldnt load the shared file", err)
        }
        if !reflect.DeepEqual(carolfilebytes, checkthefile) {
                t.Error("Incorrect load for carol", carolfilebytes, checkthefile)
        }
        bobappendbytes := []byte("Hey I am bob and i just got this file/")
        checkthefile = append(checkthefile, bobappendbytes...)
        err = bob1.AppendFile("break it bob", bobappendbytes)
        if err != nil {
                t.Error("bob failed to append to the file", err)
                return
        }
        aliceloadingbobappend, err := alice3.LoadFile("will it break")
        if err != nil {
                t.Error("Failed to load bobs append for alice", err)
                return
        }
  if !reflect.DeepEqual(aliceloadingbobappend, checkthefile) {
                t.Error("not same as it should be once appended by bob", aliceloadingbobappend, checkthefile)
                return
        }
        carolloadingbobappend, err := carol2.LoadFile("break it carol")
        if err != nil {
                t.Error("Failed to load bobs append for carol", err)
                return
        }
        if !reflect.DeepEqual(carolloadingbobappend, checkthefile) {
                t.Error("not same as it should be once appended by bob", carolloadingbobappend, checkthefile)
                return
        }
        carolappendbytes := []byte("yo yo homies i am carol and i just got this file/")
        checkthefile = append(checkthefile, carolappendbytes...)
        err = carol3.AppendFile("break it carol", carolappendbytes)
        if err != nil {
                t.Error("carol failed to append to the file", err)
                return
        }
        aliceloadingcarolappend, err := alice2.LoadFile("will it break")
        if err != nil {
                t.Error("Failed to load carols append for alice", err)
                return
        }
  if !reflect.DeepEqual(aliceloadingcarolappend, checkthefile) {
                t.Error("not same as it should be once appended by bob", aliceloadingcarolappend, checkthefile)
                return
        }
        bobloadingcarolappend, err := bob2.LoadFile("break it bob")
        if err != nil {
                t.Error("Failed to load carols append for bob", err)
                return
        }
        if !reflect.DeepEqual(bobloadingcarolappend, checkthefile) {
                t.Error("not same as it should be once appended by bob", bobloadingcarolappend, checkthefile)
                return
        } }

        //-------------------------------------------------------


        func TestAlotOfThings3(t *testing.T) {
        clear()
        _, err1 := InitUser("alice", "first")
        if err1 != nil {
                t.Error("Failed to initialize user", err1)
                return
        }
        _, err2 := InitUser("bob", "second")
        if err2 != nil {
                t.Error("Failed to initialize user", err2)
                return
        }
        _, err3 := InitUser("carol", "third")
        if err3 != nil {
                t.Error("Failed to initialize user", err3)
                return
        }
        alice1, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice2, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice3, err := GetUser("alice", "first")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob1, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob2, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        bob3, err := GetUser("bob", "second")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol1, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol2, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        carol3, err := GetUser("carol", "third")
        if err != nil {
                t.Error("Failed to get user", err)
                return
        }
        alice1file1 := []byte("this is just a file to store for alice")
        alice2file2 := []byte("this will be another file stored by alice")
        alice3file3 := []byte("yet another file that alice will store")
        bob1file1 := []byte("this is just a file to store for bob")
        bob2file2 := []byte("this will be another file stored for bob")
        bob3file3 := []byte("yet another file that bob will store")
        carol1file1 := []byte("this is just a file to store for carol")
        carol2file2 := []byte("this will be another file stored for carol")
        carol3file3 := []byte("yet another file that bob will store")
        alice1.StoreFile("file1", alice1file1)
        alice2.StoreFile("file2", alice2file2)
        alice3.StoreFile("file3", alice3file3)
        bob1.StoreFile("file1", bob1file1)
        bob2.StoreFile("file2", bob2file2)
        bob3.StoreFile("file3", bob3file3)
        carol1.StoreFile("file1", carol1file1)
        carol2.StoreFile("file2", carol2file2)
        carol3.StoreFile("file3", carol3file3)



        firstbytesforfile := []byte("The start of a shared file/")
        checkthefile := firstbytesforfile
        alice1.StoreFile("will it break", firstbytesforfile)
        secondbytesforfile := []byte("Second Instance of alice has appened to this file/")
        checkthefile = append(checkthefile, secondbytesforfile...)
        err = alice2.AppendFile("will it break", secondbytesforfile)
        if err != nil {
               t.Error("Failed to append to will it break", err)
               return
        }
        thirdbytesforfile := []byte("Third Instance of alice has appened even more/")
        checkthefile = append(checkthefile, thirdbytesforfile...)
        err = alice3.AppendFile("will it break", thirdbytesforfile)
        if err != nil {
               t.Error("Failed to append to will it break", err)
               return
        }
        alice1appendedfromload, err := alice1.LoadFile("will it break")
        if err != nil {
               t.Error("Failed to load all the appeneds", err)
        }
        if !reflect.DeepEqual(alice1appendedfromload, checkthefile) {
               t.Error("Incorrect appending", alice1appendedfromload, checkthefile)
        }

        magicstringforbob, err := alice1.ShareFile("will it break", "bob")
        if err != nil {
                t.Error("Failed to share the file with bob", err)
                return
        }
        err = bob2.ReceiveFile("break it bob", "alice", magicstringforbob)
        if err != nil {
                t.Error("received file broke for bob", err)
                return
        }
        bobfilebytes, err := bob3.LoadFile("break it bob")
        if err!=nil {
               t.Error("bob couldnt load the shared file", err)
        }
        if !reflect.DeepEqual(bobfilebytes, checkthefile) {
               t.Error("incorrect load for bob", bobfilebytes, checkthefile)
        }

        magicstringforcarol, err := alice2.ShareFile("will it break", "carol")
        if err != nil {
                t.Error("Failed to share the file with carol", err)
                return
        }
        err = carol3.ReceiveFile("break it carol", "alice", magicstringforcarol)
        if err != nil {
                t.Error("received file broke for carol", err)
                return
        }
        carolfilebytes, err := carol1.LoadFile("break it carol")
        if err!=nil {
              t.Error("carol couldnt load the shared file", err)
        }
        if !reflect.DeepEqual(carolfilebytes, checkthefile) {
               t.Error("Incorrect load for carol", carolfilebytes, checkthefile)
        }
        bobappendbytes := []byte("Hey I am bob and i just got this file/")
        checkthefile = append(checkthefile, bobappendbytes...)
        err = bob1.AppendFile("break it bob", bobappendbytes)
        if err != nil {
               t.Error("bob failed to append to the file", err)
               return
        }
        aliceloadingbobappend, err := alice3.LoadFile("will it break")
        if err != nil {
                t.Error("Failed to load bobs append for alice", err)
                return
        }
  if !reflect.DeepEqual(aliceloadingbobappend, checkthefile) {
               t.Error("not same as it should be once appended by bob", aliceloadingbobappend, checkthefile)
               return
        }
        carolloadingbobappend, err := carol2.LoadFile("break it carol")
        if err != nil {
               t.Error("Failed to load bobs append for carol", err)
               return
        }
        if !reflect.DeepEqual(carolloadingbobappend, checkthefile) {
               t.Error("not same as it should be once appended by bob", carolloadingbobappend, checkthefile)
               return
        }
        carolappendbytes := []byte("yo yo homies i am carol and i just got this file/")
        checkthefile = append(checkthefile, carolappendbytes...)
        err = carol3.AppendFile("break it carol", carolappendbytes)
        if err != nil {
               t.Error("carol failed to append to the file", err)
               return
        }
        aliceloadingcarolappend, err := alice2.LoadFile("will it break")
        if err != nil {
               t.Error("Failed to load carols append for alice", err)
               return
       }
  if !reflect.DeepEqual(aliceloadingcarolappend, checkthefile) {
               t.Error("not same as it should be once appended by bob", aliceloadingcarolappend, checkthefile)
                return
        }
        bobloadingcarolappend, err := bob2.LoadFile("break it bob")
        if err != nil {
               t.Error("Failed to load carols append for bob", err)
               return
        }
       if !reflect.DeepEqual(bobloadingcarolappend, checkthefile) {
               t.Error("not same as it should be once appended by bob", bobloadingcarolappend, checkthefile)
               return
       }


        err = alice2.RevokeFile("will it break", "carol")
        if err != nil {
                t.Error("alice could not revoke make it break from carol", err)
                return
        }
        cancarolgetcontents, err := carol3.LoadFile("break it carol")
        if err == nil {
                t.Error("carol is still able to retireve file after being revoked", cancarolgetcontents)
        }
        aliceaddingafterrevoke := []byte("hey bob carol cant see the file anymore!/")
        checkthefile = append(checkthefile, aliceaddingafterrevoke...)
        err = alice3.AppendFile("will it break", aliceaddingafterrevoke)
        if err != nil {
                t.Error("alice failed to append to the file after a revoke", err)
                return
        }
        bobloadingfromalicenewbytes, err := bob3.LoadFile("break it bob")
        if err != nil {
                t.Error("Failed to load alice append for bob after she revoked carol", err)
                return
        }
  		if !reflect.DeepEqual(bobloadingfromalicenewbytes, checkthefile) {
                t.Error("not same as it should be once appended by alice and loaded", aliceloadingbobappend, checkthefile)
                return
        }
        bobaddingafterrevoke := []byte("hey alice it so funny that carol cant see this awesome file/")
        checkthefile = append(checkthefile, bobaddingafterrevoke...)
        err = bob1.AppendFile("break it bob", bobaddingafterrevoke)
        if err != nil {
                t.Error("bob failed to append to the file after a revoke", err)
                return
        }
        aliceloadingfrombobnewbytes, err := alice1.LoadFile("will it break")
        if err != nil {
                t.Error("Failed to load bobs append for alice after she revoked carol", err)
                return
        }
  if !reflect.DeepEqual(aliceloadingfrombobnewbytes, checkthefile) {
                t.Error("not same as it should be once appended by alice and loaded", aliceloadingbobappend, checkthefile)
                return
        }
        err = alice3.RevokeFile("will it break", "bob")
        if err != nil {
                t.Error("alice could not revoke make it break from bob", err)
                return
        }
        canbobgetcontents, err := bob3.LoadFile("break it bob")
        if err == nil {
                t.Error("bob is still able to retireve file after being revoked", canbobgetcontents)
        }
        finalbytes := []byte("HAHAHA Bob you fool your in the same situation as carol!/")
        checkthefile = append(checkthefile, finalbytes...)
        err = alice2.AppendFile("will it break", finalbytes)
        if err != nil {
                t.Error("alice failed to append to the file after a the final revoke", err)
                return
        }
        alicefinalbytes, err := alice1.LoadFile("will it break")
        if err != nil {
                t.Error("Failed to load alices append for final case", err)
                return
        }
  if !reflect.DeepEqual(alicefinalbytes, checkthefile) {
                t.Error("not same as it should be once appended by alice and loaded", alicefinalbytes, checkthefile)
                return
        }
}
