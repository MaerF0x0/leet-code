/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // Deal with empty lists cause i'll probably end up with nil pointer panics if I dont
    if list1 == nil { 
        return list2
    } else if list2 == nil {
        return list1
    }
    
    
    // Pick a node to be the head of our new list
    var head *ListNode
    if list1.Val <= list2.Val {
        head = list1
        list1 = list1.Next
    } else {
        head = list2
        list2 = list2.Next
    }
    
    
    // Now iterate the 2 lists 
    cur := head
    for {
// See if we should ffw remainder and break
        if list1 == nil {
            cur.Next = list2
            break
        } else if list2 == nil {
            cur.Next = list1
            break
        }        

        if list1.Val <= list2.Val {
            cur.Next = list1
            list1 = list1.Next
        } else {
            cur.Next = list2
            list2 = list2.Next
        }
        
        cur = cur.Next

        
    }
    
    return head
}