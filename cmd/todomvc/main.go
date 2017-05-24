package main

import (
	"fmt"
	"github.com/dtylman/gowd"
	"os"
)

const (
	filterAll       = "ShowAll"
	filterActive    = "ShowActive"
	filterCompleted = "ShowCompleted"
)

type todo struct {
	completed bool
	text      string
}

type app struct {
	*gowd.Element
	inputTodo         *gowd.Element
	todoList          *gowd.Element
	todocount         *gowd.Element
	btnClearCompleted *gowd.Element
	filter            string
	todos             []todo
}

func newApplication() *app {
	a := &app{}
	a.todos = make([]todo, 0)
	a.Element = gowd.NewElement("section")
	a.SetClass("todoapp")
	em := gowd.NewElementMap()

	header := `<header class="header">
				<h1>todos</h1>
				<input class="new-todo" placeholder="What needs to be done?" id="inputTodo" autofocus>
			</header>`
	a.AddHTML(header, em)
	a.inputTodo = em["inputTodo"]
	a.inputTodo.OnKeyPressEvent(gowd.OnKeyPress, 13, a.inputTodoEnter)

	main := `<section class="main">
				<input class="toggle-all" id="toggle-all" type="checkbox">
				<label for="toggle-all">Mark all as complete</label>
				<ul class="todo-list" id="todoList">
				</ul>
			</section>`
	a.AddHTML(main, em)
	a.todoList = em["todoList"]
	em["toggle-all"].OnEvent(gowd.OnClick, a.btnToggleAllClicked)

	footer := `<footer class="footer">
				<span class="todo-count" id="todocount"></span>
				<ul class="filters">
					<li>
						<a class="selected" href="#/" id="btnShowAll">All</a>
					</li>
					<li>
						<a href="#/active" id="btnShowActive">Active</a>
					</li>
					<li>
						<a href="#/completed" id="btnShowCompleted">Completed</a>
					</li>
				</ul>
				<button class="clear-completed" id="clearCompleted" >Clear completed</button>
			</footer>`
	a.AddHTML(footer, em)
	a.todocount = em["todocount"]
	a.btnClearCompleted = em["clearCompleted"]
	a.btnClearCompleted.OnEvent(gowd.OnClick, a.btnClearCompletedClicked)
	em["btnShowAll"].OnEvent(gowd.OnClick, a.btnFilterClicked)
	em["btnShowAll"].Object = filterAll
	em["btnShowActive"].OnEvent(gowd.OnClick, a.btnFilterClicked)
	em["btnShowActive"].Object = filterActive
	em["btnShowCompleted"].OnEvent(gowd.OnClick, a.btnFilterClicked)
	em["btnShowCompleted"].Object = filterCompleted
	a.updateList()
	return a
}

func (a *app) updateList() {
	a.todoList.RemoveElements()
	completed := 0
	for i, t := range a.todos {
		if a.filter == filterCompleted && !t.completed {
			continue
		} else if a.filter == filterActive && t.completed {
			continue
		}
		li := gowd.NewElement("li")
		if t.completed {
			li.SetClass("completed")
			completed++
		}

		div := gowd.NewElement("div")
		div.SetClass("view")
		li.AddElement(div)

		toggle := gowd.NewElement("input")
		toggle.Object = i
		toggle.SetClass("toggle")
		toggle.SetAttribute("type", "checkbox")
		toggle.OnEvent(gowd.OnClick, a.btnToggleClicked)
		if t.completed {
			toggle.SetAttribute("checked", "")
		}
		div.AddElement(toggle)

		label := gowd.NewElement("label")
		label.SetText(t.text)
		div.AddElement(label)

		button := gowd.NewElement("button")
		button.SetClass("destroy")
		button.Object = i
		button.OnEvent(gowd.OnClick, a.btnRemoveClicked)
		div.AddElement(button)

		edit := gowd.NewElement("input")
		edit.SetClass("edit")
		edit.SetAttribute("value", t.text)
		li.AddElement(edit)

		a.todoList.AddElement(li)
	}
	if completed == 0 {
		a.btnClearCompleted.Hide()
	} else {
		a.btnClearCompleted.Show()
	}
	total := len(a.todos)
	item := "item"
	if total != 1 {
		item = item + "s"
	}
	a.todocount.RemoveElements()
	a.todocount.AddHTML(fmt.Sprintf(`<strong>%d </strong> %s left`, total, item), nil)
}

func (a *app) btnRemoveClicked(sender *gowd.Element, event *gowd.EventElement) {
	index := sender.Object.(int)
	a.todos = append(a.todos[:index], a.todos[index+1:]...)
	a.updateList()
}

func (a *app) inputTodoEnter(sender *gowd.Element, event *gowd.EventElement) {
	if sender.GetValue() == "" {
		return
	}
	a.todos = append(a.todos, todo{text: sender.GetValue(), completed: false})
	sender.SetAttribute("value", "")
	a.updateList()
}

func (a *app) btnToggleClicked(sender *gowd.Element, event *gowd.EventElement) {
	index := sender.Object.(int)
	_, checked := sender.GetAttribute("checked")
	a.todos[index].completed = checked
	a.updateList()
}

func (a *app) btnToggleAllClicked(sender *gowd.Element, event *gowd.EventElement) {
	_, checked := sender.GetAttribute("checked")
	for i := range a.todos {
		a.todos[i].completed = checked
	}
	a.updateList()
}

func (a *app) btnClearCompletedClicked(sender *gowd.Element, event *gowd.EventElement) {
	newList := make([]todo, 0)
	for i := range a.todos {
		if !a.todos[i].completed {
			newList = append(newList, a.todos[i])
		}
	}
	a.todos = newList
	a.updateList()
}

func (a *app) btnFilterClicked(sender *gowd.Element, event *gowd.EventElement) {
	a.filter = sender.Object.(string)
	a.Find("btnShowAll").UnsetClass("selected")
	a.Find("btnShowActive").UnsetClass("selected")
	a.Find("btnShowCompleted").UnsetClass("selected")
	sender.SetClass("selected")
	a.updateList()
}

func main() {
	err := gowd.Run(newApplication().Element)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
