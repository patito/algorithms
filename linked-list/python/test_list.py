import unittest

from linked_list import LinkedList


class TestLinkedList(unittest.TestCase):

    def test_append(self):
        l = LinkedList()

        l.append("hobbit")
        self.assertEqual(l.head.data, "hobbit")
        self.assertEqual(l.last.data, "hobbit")
        self.assertEqual(len(l), 1)

        l.append("dwarf")
        self.assertEqual(l.head.data, "hobbit")
        self.assertEqual(l.last.data, "dwarf")
        self.assertEqual(len(l), 2)

        l.append("say")
        l.append("my")
        l.append("name")
        self.assertEqual(l.head.data, "hobbit")
        self.assertEqual(l.last.data, "name")
        self.assertEqual(len(l), 5)

    def test_preppend(self):
        l = LinkedList()

        l.preppend("dead")
        self.assertEqual(l.head.data, "dead")
        self.assertEqual(l.last.data, "dead")
        self.assertEqual(len(l), 1)

        l.preppend("finger")
        self.assertEqual(l.head.data, "finger")
        self.assertEqual(l.last.data, "dead")
        self.assertEqual(len(l), 2)

        l.preppend("do")
        l.preppend("not")
        l.preppend("code")
        self.assertEqual(l.head.data, "code")
        self.assertEqual(l.last.data, "dead")
        self.assertEqual(len(l), 5)

    def test_insert(self):
        l = LinkedList()

        l.insert("WILSONNNNN", 10)
        self.assertEqual(l.head.data, "WILSONNNNN")
        self.assertEqual(l.last.data, "WILSONNNNN")
        self.assertEqual(len(l), 1)

        l.insert("douchebag", -10)
        self.assertEqual(l.head.data, "douchebag")
        self.assertEqual(l.last.data, "WILSONNNNN")
        self.assertEqual(len(l), 2)

        l.insert("Isso eh pra vc aprender a nunca mais me esnobar", 1)
        self.assertEqual(
            l.head.next.data,
            "Isso eh pra vc aprender a nunca mais me esnobar")
        self.assertEqual(len(l), 3)

    def test_is_empty(self):
        l = LinkedList()

        self.assertEqual(l.is_empty(), True)

        l.append("I don't wanna close my eyes...")
        self.assertEqual(l.is_empty(), False)

    def test_search(self):
        l = LinkedList()

        self.assertEqual(l.search("I love bugs"), False)

        l.preppend("I like to dance rihana")
        self.assertEqual(l.search("I like to dance rihana"), True)

    def test_remove(self):
        l = LinkedList()

        l.remove("remove empty")

        l.append("show me the moneyyyyyy")
        l.preppend("bazinga")
        l.insert("crazy life", 1)
        l.remove("crazy life")
        self.assertEqual(l.search("crazy life"), False)
        self.assertEqual(l.search("show me the moneyyyyyy"), True)
        self.assertEqual(l.search("bazinga"), True)
        self.assertEqual(len(l), 2)

        l.remove("show me the moneyyyyyy")
        l.remove("bazinga")
        self.assertEqual(l.search("show me the moneyyyyyy"), False)
        self.assertEqual(l.search("bazinga"), False)
        self.assertEqual(len(l), 0)

    def test_pop(self):
        l = LinkedList()

        l.pop("remove empty")

        l.append("I Was a Mermaid")
        l.preppend("and now")
        l.insert("I'm a Pop Star", 1)
        self.assertEqual(len(l), 3)
        self.assertEqual(l.search("and now"), True)

        l.pop(0)
        self.assertEqual(l.search("and now"), False)
        self.assertEqual(len(l), 2)

        l.pop(1)
        self.assertEqual(l.search("I Was a Mermaid"), False)
        self.assertEqual(len(l), 1)

        l.pop(0)
        self.assertEqual(l.search("I'm a Pop Star"), False)
        self.assertEqual(len(l), 0)


if __name__ == '__main__':
    unittest.main()
