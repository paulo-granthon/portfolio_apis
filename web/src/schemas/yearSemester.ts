export class YearSemester {
  public year: number;
  public semester: 1 | 2;

  constructor(year: number, semester: 1 | 2) {
    this.year = year;
    this.semester = semester;
  }

  static now() {
    const now = new Date();
    return new YearSemester(now.getFullYear(), now.getMonth() < 6 ? 1 : 2);
  }

  // add one semester to YearSemester
  plusSemesters(semesters: number) {
    const newSemester = this.semester - 1 + semesters;
    return new YearSemester(
      this.year + Math.floor(newSemester / 2),
      ((newSemester % 2) + 1) as 1 | 2,
    );
  }

  toString() {
    return `${this.year}-${this.semester}`;
  }

  // implicit conversion from { year: number, semester: 1 | 2 } to YearSemester
  static fromObject({ year, semester }: { year: number; semester: 1 | 2 }) {
    return new YearSemester(year, semester);
  }

  // function that returns the number of semesters between this and another YearSemester
  semestersBetween(other: YearSemester) {
    return (other.year - this.year) * 2 + (other.semester - this.semester);
  }

  // function that returns what's the current semester of someone that matriculated in this YearSemester
  currentSemester() {
    const now = YearSemester.now().plusSemesters(1);
    const semesters = this.semestersBetween(now);
    return semesters < 0 ? 0 : semesters;
  }
}
