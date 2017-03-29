package bootstrap

import (
	"github.com/dtylman/gowd"
)

/*
<nav class="navbar navbar-default">
  <div class="container-fluid">
      <ul class="nav navbar-nav">
        <button type="button" class="btn btn-default navbar-btn">Sign in</button>
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
          <ul class="dropdown-menu">
            <li><a href="#">Action</a></li>
            <li><a href="#">Another action</a></li>
            <li><a href="#">Something else here</a></li>
            <li role="separator" class="divider"></li>
            <li><a href="#">Separated link</a></li>
          </ul>
        </li>
      </ul>

  </div><!-- /.container-fluid -->
</nav>

*/
type Navbar struct {
	*gowd.Element
	ul          *gowd.Element
	navElements []*gowd.Element
}

func NewNavBar() *Navbar {
	nb := &Navbar{}
	nb.ul = NewElement("ul", "nav navbar-nav")
	nb.Element = NewElement("nav", "navbar navbar-default")
	div := NewElement("div", "container-fluid")
	div.AddElement(nb.ul)
	nb.Element.AddElement(div)
	nb.navElements = make([]*gowd.Element, 0)
	return nb
}

func (nb *Navbar) AddNavElement(elem *gowd.Element) {
	li := gowd.NewElement("li")
	li.AddElement(elem)
	nb.ul.AddElement(li)
	nb.navElements = append(nb.navElements, elem)
}

func (nb *Navbar) AddLinkButton(caption string) *gowd.Element {
	btn := NewLinkButton(caption)
	nb.AddNavElement(btn)
	return btn
}

func (nb *Navbar) AddButton(buttonType string, caption string) *gowd.Element {
	btn := NewButton(buttonType + " navbar-btn", caption)
	nb.AddNavElement(btn)
	return btn
}

func (nb *Navbar) AddSeperator() {
	li := NewElement("li", "divider")
	li.SetAttribute("role", "seperator")
	nb.ul.AddElement(li)
}

func (nb*Navbar) SetActiveElement(elem*gowd.Element) {
	for i := range nb.navElements {
		nb.navElements[i].UnsetClass("active")
	}
	elem.SetClass("active")

}