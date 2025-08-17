export function generateTimeSlots(
  startHour = 9,
  endHour = 18,
  slotMinutes = 60
) {
  const slots: { start: string; end: string }[] = [];
  for (let h = startHour; h < endHour; h++) {
    const start = `${String(h).padStart(2, "0")}:00`;
    const end = `${String(h + 1).padStart(2, "0")}:00`;
    slots.push({ start, end });
  }
  return slots;
}