package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"sort"
	"strconv"
	"strings"
)

func entry(value string) *widget.Entry { e := widget.NewEntry(); e.SetText(value); return e }
func label(s string) *widget.Label     { l := widget.NewLabel(s); l.Wrapping = fyne.TextWrapWord; return l }
func task(title, description string, inputs fyne.CanvasObject, run func() (string, error)) fyne.CanvasObject {
	out := label("Результат появится здесь.")
	button := widget.NewButton("Выполнить", func() {
		s, e := run()
		if e != nil {
			out.SetText("Ошибка: " + e.Error())
			return
		}
		out.SetText(s)
	})
	return widget.NewCard(title, "", container.NewVBox(label(description), inputs, button, out))
}
func integer(e *widget.Entry) (int, error) {
	v, err := strconv.Atoi(strings.TrimSpace(e.Text))
	if err != nil || v < -1000000 || v > 1000000 {
		return 0, fmt.Errorf("введите целое число от −1000000 до 1000000")
	}
	return v, nil
}
func number(e *widget.Entry) (float64, error) {
	v, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(e.Text), ",", "."), 64)
	if err != nil || math.IsNaN(v) || math.IsInf(v, 0) {
		return 0, fmt.Errorf("введите конечное число")
	}
	return v, nil
}
func form(items ...*widget.FormItem) fyne.CanvasObject       { return widget.NewForm(items...) }
func item(name string, o fyne.CanvasObject) *widget.FormItem { return widget.NewFormItem(name, o) }
func page(objects ...fyne.CanvasObject) fyne.CanvasObject {
	return container.NewVScroll(container.NewVBox(objects...))
}

func chapter2() fyne.CanvasObject {
	a, b := entry("12"), entry("3")
	op := widget.NewSelect([]string{"+", "−", "×", "÷"}, nil)
	op.SetSelected("+")
	calc := task("1. Калькулятор", "Арифметика, switch. Дробную часть можно вводить через точку или запятую.", form(item("A", a), item("Операция", op), item("B", b)), func() (string, error) {
		x, e := number(a)
		if e != nil {
			return "", e
		}
		y, e := number(b)
		if e != nil {
			return "", e
		}
		r, e := Calculate(x, y, op.Selected)
		return fmt.Sprintf("%g %s %g = %g", x, op.Selected, y, r), e
	})
	n := entry("7")
	check := task("2. Проверка числа", "Определите знак и чётность целого числа с помощью if.", form(item("Число", n)), func() (string, error) {
		v, e := integer(n)
		if e != nil {
			return "", e
		}
		sign := "ноль"
		if v > 0 {
			sign = "положительное"
		} else if v < 0 {
			sign = "отрицательное"
		}
		parity := "чётное"
		if v%2 != 0 {
			parity = "нечётное"
		}
		return fmt.Sprintf("%d: %s, %s", v, sign, parity), nil
	})
	m := entry("5")
	table := task("3. Таблица умножения", "Цикл for: множители от 1 до 10.", form(item("Число", m)), func() (string, error) {
		v, e := integer(m)
		if e != nil {
			return "", e
		}
		var s strings.Builder
		for i := 1; i <= 10; i++ {
			fmt.Fprintf(&s, "%d × %d = %d\n", v, i, v*i)
		}
		return strings.TrimSpace(s.String()), nil
	})
	arr := entry("5, 2, 9, -1, 4")
	array := task("4. Массив из пяти чисел", "Заполните настоящий [5]int. Найдите сумму, минимум, максимум и среднее.", form(item("Ровно 5 чисел", arr)), func() (string, error) {
		v, e := ParseNumbers(arr.Text)
		if e != nil {
			return "", e
		}
		if len(v) != 5 {
			return "", fmt.Errorf("нужно ровно пять чисел")
		}
		var fixed [5]int
		copy(fixed[:], v)
		lo, hi := MinMax(fixed[:])
		return fmt.Sprintf("Массив: %v\nСумма: %d; min: %d; max: %d; среднее: %.2f", fixed, Sum(fixed[:]...), lo, hi, float64(Sum(fixed[:]...))/5), nil
	})
	return page(calc, check, table, array)
}

func chapter3() fyne.CanvasObject {
	nums := entry("1 2 3 4")
	sum := task("1. Функции и результаты", "Sum(values ...int) и MinMax возвращают результаты. Срез передаётся с ...", form(item("Числа", nums)), func() (string, error) {
		v, e := ParseNumbers(nums.Text)
		if e != nil {
			return "", e
		}
		lo, hi := MinMax(v)
		return fmt.Sprintf("Сумма: %d; минимум: %d; максимум: %d", Sum(v...), lo, hi), nil
	})
	n := entry("5")
	fact := task("2. Рекурсивный факториал", "Базовый случай 0! = 1. Допустимый диапазон 0–20 предотвращает переполнение uint64.", form(item("N", n)), func() (string, error) {
		v, e := integer(n)
		if e != nil {
			return "", e
		}
		if v < 0 || v > 20 {
			return "", fmt.Errorf("нужно N от 0 до 20")
		}
		return fmt.Sprintf("%d! = %d", v, Factorial(v)), nil
	})
	values := entry("2 3 4")
	choice := widget.NewSelect([]string{"Квадрат", "Удвоение"}, nil)
	choice.SetSelected("Квадрат")
	transform := task("3. Функция как параметр", "Transform получает анонимную функцию и применяет её к каждому числу.", form(item("Числа", values), item("Преобразование", choice)), func() (string, error) {
		v, e := ParseNumbers(values.Text)
		if e != nil {
			return "", e
		}
		f := func(x int) int { return x * x }
		if choice.Selected == "Удвоение" {
			f = func(x int) int { return x * 2 }
		}
		return fmt.Sprintf("%v → %v", v, Transform(v, f)), nil
	})
	next := Counter()
	out := label("Счётчик: 0")
	counter := widget.NewCard("4. Замыкание", "", container.NewVBox(label("Функция хранит n между вызовами. При перезапуске приложения счётчик сбрасывается."), container.NewHBox(widget.NewButton("Следующее", func() { out.SetText(fmt.Sprintf("Счётчик: %d", next())) }), widget.NewButton("Сбросить", func() { next = Counter(); out.SetText("Счётчик: 0") })), out))
	return page(sum, fact, transform, counter)
}

func chapter4() fyne.CanvasObject {
	n := entry("10")
	inc := task("1. Значение и указатель", "Сравните передачу копии и изменение исходной переменной через *int.", form(item("Исходное число", n)), func() (string, error) {
		v, e := integer(n)
		if e != nil {
			return "", e
		}
		copyResult := IncrementValue(v)
		before := v
		IncrementPointer(&v)
		return fmt.Sprintf("По значению: функция вернула %d, исходная переменная осталась %d.\nПо указателю: исходная переменная стала %d.\nАдрес: %p; значение *p: %d", copyResult, before, v, &v, v), nil
	})
	a, b := entry("3"), entry("8")
	swap := task("2. Обмен через указатели", "Swap(&a, &b) меняет обе переменные.", form(item("A", a), item("B", b)), func() (string, error) {
		x, e := integer(a)
		if e != nil {
			return "", e
		}
		y, e := integer(b)
		if e != nil {
			return "", e
		}
		before := fmt.Sprintf("До: A=%d, B=%d", x, y)
		Swap(&x, &y)
		return fmt.Sprintf("%s\nПосле: A=%d, B=%d", before, x, y), nil
	})
	return page(inc, swap)
}

func chapter5() fyne.CanvasObject {
	name, age, city := entry("Анна"), entry("20"), entry("Кызылорда")
	student := task("1. Структура и методы", "Student содержит Address; Age — псевдоним int. Метод Birthday имеет получатель *Student.", form(item("Имя", name), item("Возраст", age), item("Город", city)), func() (string, error) {
		v, e := integer(age)
		if e != nil {
			return "", e
		}
		if v < 0 || v > 150 {
			return "", fmt.Errorf("возраст должен быть от 0 до 150")
		}
		if strings.TrimSpace(name.Text) == "" || strings.TrimSpace(city.Text) == "" {
			return "", fmt.Errorf("заполните имя и город")
		}
		s := Student{strings.TrimSpace(name.Text), v, Address{strings.TrimSpace(city.Text)}}
		before := s.Description()
		s.Birthday()
		return "Создан: " + before + "\nПосле Birthday(): " + s.Description(), nil
	})
	nums, extra := entry("8 3 5"), entry("2")
	slice := task("2. Срез: append и сортировка", "Добавьте число, отсортируйте копию среза. Покажите len и cap.", form(item("Числа", nums), item("Добавить", extra)), func() (string, error) {
		v, e := ParseNumbers(nums.Text)
		if e != nil {
			return "", e
		}
		x, e := integer(extra)
		if e != nil {
			return "", e
		}
		v = append(v, x)
		sorted := append([]int(nil), v...)
		sort.Ints(sorted)
		return fmt.Sprintf("После append: %v\nlen=%d, cap=%d\nСортированная копия: %v", v, len(v), cap(v), sorted), nil
	})
	words := entry("go код go функции код go")
	freq := task("3. Карта частот слов", "map[string]int: посчитайте слова без учёта регистра. Разделитель — пробел; знаки препинания остаются частью слова.", form(item("Текст", words)), func() (string, error) {
		v := strings.Fields(strings.ToLower(words.Text))
		if len(v) == 0 || len(v) > 1000 {
			return "", fmt.Errorf("введите от 1 до 1000 слов")
		}
		counts := map[string]int{}
		for _, w := range v {
			counts[w]++
		}
		keys := make([]string, 0, len(counts))
		for k := range counts {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var s strings.Builder
		for _, k := range keys {
			fmt.Fprintf(&s, "%s: %d\n", k, counts[k])
		}
		return strings.TrimSpace(s.String()), nil
	})
	contacts := map[string]string{}
	key, value := entry("Анна"), entry("+7 700 000 00 00")
	out := label("Справочник пуст.")
	show := func() {
		keys := make([]string, 0, len(contacts))
		for k := range contacts {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var s strings.Builder
		for _, k := range keys {
			fmt.Fprintf(&s, "%s: %s\n", k, contacts[k])
		}
		if len(keys) == 0 {
			out.SetText("Справочник пуст.")
		} else {
			out.SetText(strings.TrimSpace(s.String()))
		}
	}
	directory := widget.NewCard("4. Карта: справочник", "", container.NewVBox(label("Добавление, поиск через value, ok и delete. Данные хранятся до закрытия приложения."), form(item("Имя (ключ)", key), item("Телефон", value)), container.NewHBox(widget.NewButton("Сохранить", func() {
		k, v := strings.TrimSpace(key.Text), strings.TrimSpace(value.Text)
		if k == "" || v == "" {
			out.SetText("Ошибка: заполните имя и телефон.")
			return
		}
		contacts[k] = v
		show()
	}), widget.NewButton("Найти", func() {
		k := strings.TrimSpace(key.Text)
		v, ok := contacts[k]
		if !ok {
			out.SetText("Ключ не найден.")
		} else {
			out.SetText(k + ": " + v)
		}
	}), widget.NewButton("Удалить", func() { delete(contacts, strings.TrimSpace(key.Text)); show() }), widget.NewButton("Показать все", show)), out))
	return page(student, slice, freq, directory)
}

func buildUI() *fyne.Container {
	tabs := container.NewAppTabs(
		container.NewTabItem("Глава 2 · Основы", chapter2()),
		container.NewTabItem("Глава 3 · Функции", chapter3()),
		container.NewTabItem("Глава 4 · Указатели", chapter4()),
		container.NewTabItem("Глава 5 · Типы", chapter5()),
	)
	return container.NewStack(tabs)
}