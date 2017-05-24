package bootstrap

import "github.com/dtylman/gowd"

/*

<nav aria-label="Page navigation">
  <ul class="pagination">
    <li>
      <a href="#" aria-label="Previous">
        <span aria-hidden="true">&laquo;</span>
      </a>
    </li>
    <li><a href="#">1</a></li>
    <li><a href="#">2</a></li>
    <li><a href="#">3</a></li>
    <li><a href="#">4</a></li>
    <li><a href="#">5</a></li>
    <li>
      <a href="#" aria-label="Next">
        <span aria-hidden="true">&raquo;</span>
      </a>
    </li>
  </ul>
</nav>
*/

type Pagination struct {
	*gowd.Element
	Items *List
}

func NewPagination() *Pagination {
	p := new(Pagination)
	p.Element = gowd.NewElement("nav")
	p.Items = NewList(ListUnordered, "pagination")
	p.AddElement(p.Items.Element)
	return p

}

func (p *Pagination) AddItem(caption string, active bool, handler gowd.EventHandler) *gowd.Element {
	link := NewLinkButton(caption)
	if handler != nil {
		link.OnEvent(gowd.OnClick, handler)
	}
	item := p.Items.AddItem(link)
	if active {
		item.SetClass("active")
	}
	return link
}
